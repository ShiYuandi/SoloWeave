package installer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ShiYuandi/SoloWeave/internal/bundle"
	"github.com/ShiYuandi/SoloWeave/internal/project"
)

const Version = "0.1.0-preview"

type Change struct {
	Path   string
	Action string
}
type Manifest struct {
	Version string            `json:"version"`
	Agents  []string          `json:"agents"`
	Files   map[string]string `json:"files"`
}

func manifestPath(root string) string { return filepath.Join(root, ".soloweave", "installation.json") }
func digest(data []byte) string       { sum := sha256.Sum256(data); return hex.EncodeToString(sum[:]) }

func readManifest(root string) (Manifest, error) {
	path, err := safeTarget(root, ".soloweave/installation.json")
	if err != nil {
		return Manifest{}, err
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return Manifest{Files: map[string]string{}}, nil
	}
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return Manifest{}, fmt.Errorf("invalid installation manifest: %w", err)
	}
	if m.Files == nil {
		return Manifest{}, errors.New("invalid installation manifest: missing files")
	}
	return m, nil
}

func safeTarget(root, rel string) (string, error) {
	if filepath.IsAbs(rel) || rel == ".." || strings.HasPrefix(filepath.Clean(rel), ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("unsafe install path: %s", rel)
	}
	root, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	path := root
	for _, part := range strings.Split(filepath.Clean(rel), string(filepath.Separator)) {
		path = filepath.Join(path, part)
		info, err := os.Lstat(path)
		if err == nil && info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("symlink in install path: %s", path)
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
	}
	return path, nil
}

func addSkills(files map[string][]byte, target string) error {
	return fs.WalkDir(bundle.Files, "assets/skills", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		data, err := bundle.Files.ReadFile(path)
		if err != nil {
			return err
		}
		rel := strings.TrimPrefix(path, "assets/skills/")
		files[filepath.ToSlash(filepath.Join(target, rel))] = data
		return nil
	})
}

func ruleBytes(root, rel, asset string, existingManifest Manifest) ([]byte, error) {
	block, err := bundle.Files.ReadFile(asset)
	if err != nil {
		return nil, err
	}
	path, err := safeTarget(root, filepath.FromSlash(rel))
	if err != nil {
		return nil, err
	}
	current, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return block, nil
	}
	if err != nil {
		return nil, err
	}
	if _, managed := existingManifest.Files[rel]; !managed {
		if strings.Contains(string(current), "<!-- soloweave:start -->") {
			return nil, fmt.Errorf("existing SoloWeave block is not managed: %s", rel)
		}
		separator := "\n"
		if len(current) == 0 || current[len(current)-1] == '\n' {
			separator = ""
		}
		return append(append([]byte{}, current...), []byte(separator+"\n"+string(block))...), nil
	}
	start := strings.Index(string(current), "<!-- soloweave:start -->")
	end := strings.Index(string(current), "<!-- soloweave:end -->")
	if start < 0 || end < start {
		return nil, fmt.Errorf("managed rule block missing: %s", rel)
	}
	end += len("<!-- soloweave:end -->")
	return []byte(string(current[:start]) + strings.TrimRight(string(block), "\r\n") + string(current[end:])), nil
}

func normalizeAgents(agents []string) ([]string, error) {
	seen := map[string]bool{}
	for _, a := range agents {
		a = strings.TrimSpace(a)
		if a != "codex" && a != "claude" && a != "cursor" {
			return nil, fmt.Errorf("unknown agent: %s", a)
		}
		seen[a] = true
	}
	if len(seen) == 0 {
		return nil, errors.New("no agents selected")
	}
	var sorted []string
	for a := range seen {
		sorted = append(sorted, a)
	}
	sort.Strings(sorted)
	return sorted, nil
}

func desired(root string, agents []string, old Manifest) (map[string][]byte, error) {
	want := map[string][]byte{}
	shared := false
	for _, a := range agents {
		switch a {
		case "codex":
			shared = true
		case "cursor":
			shared = true
			data, err := bundle.Files.ReadFile("assets/rules/cursor.mdc")
			if err != nil {
				return nil, err
			}
			want[".cursor/rules/soloweave.mdc"] = data
		case "claude":
			if err := addSkills(want, ".claude/skills"); err != nil {
				return nil, err
			}
		}
	}
	if shared {
		if err := addSkills(want, ".agents/skills"); err != nil {
			return nil, err
		}
	}
	for _, spec := range []struct{ agent, path, asset string }{
		{"codex", "AGENTS.md", "assets/rules/agents.md"},
		{"claude", "CLAUDE.md", "assets/rules/claude.md"},
	} {
		for _, a := range agents {
			if a == spec.agent {
				data, err := ruleBytes(root, spec.path, spec.asset, old)
				if err != nil {
					return nil, err
				}
				want[spec.path] = data
			}
		}
	}
	return want, nil
}

func Install(root string, requested []string, dryRun bool) ([]Change, error) {
	if err := project.CheckApproval(root); err != nil {
		return nil, err
	}
	agents, err := normalizeAgents(requested)
	if err != nil {
		return nil, err
	}
	old, err := readManifest(root)
	if err != nil {
		return nil, err
	}
	agents, err = normalizeAgents(append(agents, old.Agents...))
	if err != nil {
		return nil, err
	}
	want, err := desired(root, agents, old)
	if err != nil {
		return nil, err
	}
	paths := make([]string, 0, len(want))
	for rel := range want {
		paths = append(paths, rel)
	}
	sort.Strings(paths)
	var changes []Change
	previous := map[string][]byte{}
	for _, rel := range paths {
		path, err := safeTarget(root, filepath.FromSlash(rel))
		if err != nil {
			return nil, err
		}
		current, err := os.ReadFile(path)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
		if err == nil {
			if oldHash, managed := old.Files[rel]; managed {
				if digest(current) != oldHash {
					return nil, fmt.Errorf("modified managed file: %s", rel)
				}
			} else if rel != "AGENTS.md" && rel != "CLAUDE.md" {
				return nil, fmt.Errorf("existing user file: %s", rel)
			}
			previous[rel] = current
		}
		if err != nil || digest(current) != digest(want[rel]) {
			changes = append(changes, Change{Path: rel, Action: "write"})
		}
	}
	if dryRun {
		return changes, nil
	}
	if len(changes) == 0 && old.Version == Version && strings.Join(old.Agents, ",") == strings.Join(agents, ",") {
		return nil, nil
	}
	newManifest := Manifest{Version: Version, Agents: agents, Files: map[string]string{}}
	for key, value := range old.Files {
		newManifest.Files[key] = value
	}
	var written []string
	rollback := func() {
		for i := len(written) - 1; i >= 0; i-- {
			rel := written[i]
			path := filepath.Join(root, filepath.FromSlash(rel))
			if data, existed := previous[rel]; existed {
				_ = os.WriteFile(path, data, 0644)
			} else {
				_ = os.Remove(path)
			}
		}
	}
	for _, change := range changes {
		path, err := safeTarget(root, filepath.FromSlash(change.Path))
		if err != nil {
			rollback()
			return nil, err
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			rollback()
			return nil, err
		}
		if err := os.WriteFile(path, want[change.Path], 0644); err != nil {
			rollback()
			return nil, err
		}
		written = append(written, change.Path)
	}
	for _, rel := range paths {
		newManifest.Files[rel] = digest(want[rel])
	}
	data, err := json.MarshalIndent(newManifest, "", "  ")
	if err != nil {
		rollback()
		return nil, err
	}
	manifest, err := safeTarget(root, ".soloweave/installation.json")
	if err != nil {
		rollback()
		return nil, err
	}
	if err := os.WriteFile(manifest, append(data, '\n'), 0644); err != nil {
		rollback()
		return nil, err
	}
	return changes, nil
}

func Check(root string) []string {
	m, err := readManifest(root)
	if err != nil {
		return []string{err.Error()}
	}
	if m.Version == "" {
		return []string{"installation manifest missing"}
	}
	var problems []string
	agents, err := normalizeAgents(m.Agents)
	if err != nil {
		problems = append(problems, err.Error())
	} else if want, err := desired(root, agents, m); err != nil {
		problems = append(problems, err.Error())
	} else {
		for rel := range want {
			if _, ok := m.Files[rel]; !ok {
				problems = append(problems, "untracked installed file "+rel)
			}
		}
	}
	for rel, expected := range m.Files {
		path, err := safeTarget(root, filepath.FromSlash(rel))
		if err != nil {
			problems = append(problems, err.Error())
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			problems = append(problems, "missing "+rel)
			continue
		}
		if digest(data) != expected {
			problems = append(problems, "modified "+rel)
		}
	}
	sort.Strings(problems)
	return problems
}
