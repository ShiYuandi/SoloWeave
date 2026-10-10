package continuity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Input struct {
	Task         string
	Summary      string
	Next         string
	Verification []string
	Issues       []string
}

type GitState struct {
	Available      bool     `json:"available"`
	Branch         string   `json:"branch,omitempty"`
	Head           string   `json:"head,omitempty"`
	Digest         string   `json:"digest,omitempty"`
	SnapshotDigest string   `json:"snapshot_digest,omitempty"`
	Changes        []string `json:"changes,omitempty"`
}

type Metadata struct {
	At  time.Time `json:"at"`
	Git GitState  `json:"git"`
}

func dir(root string) string { return filepath.Join(root, ".soloweave", "context") }

func safeDir(root string) error {
	for _, path := range []string{filepath.Join(root, ".soloweave"), dir(root)} {
		if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("context path is symlink: %s", path)
		}
	}
	return os.MkdirAll(dir(root), 0755)
}

func writeManaged(path string, data []byte, onlyIfMissing bool) error {
	if info, err := os.Lstat(path); err == nil {
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("context file is symlink: %s", path)
		}
		if onlyIfMissing {
			return nil
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Initialize(root, name string) error {
	if err := safeDir(root); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dir(root), "tasks"), 0755); err != nil {
		return err
	}
	files := map[string]string{
		"PROJECT.md": "# Project\n\nName: " + name + "\n\nDescribe the goal, main modules, run commands, and key entry files.\n",
		"STATUS.md":  "# Status\n\n## Completed\n\n## In Progress\n\n## Blocked\n\n## Planned\n",
		"HANDOFF.md": "# Project Handoff\n\n## Current Task\n\nNo checkpoint yet.\n\n## Next Steps\n\nCreate a task checkpoint.\n",
		"CHANGES.md": "# Changes\n\nRecord completed, important code changes.\n",
	}
	for name, body := range files {
		if err := writeManaged(filepath.Join(dir(root), name), []byte(body), true); err != nil {
			return err
		}
	}
	return nil
}

func Git(root string) GitState {
	cmd := exec.Command("git", "-C", root, "rev-parse", "--is-inside-work-tree")
	if out, err := cmd.Output(); err != nil || strings.TrimSpace(string(out)) != "true" {
		return GitState{}
	}
	state := GitState{Available: true}
	if out, err := exec.Command("git", "-C", root, "branch", "--show-current").Output(); err == nil {
		state.Branch = strings.TrimSpace(string(out))
	}
	if state.Branch == "" {
		state.Branch = "(detached)"
	}
	if out, err := exec.Command("git", "-C", root, "rev-parse", "HEAD").Output(); err == nil {
		state.Head = strings.TrimSpace(string(out))
	} else {
		state.Head = "(no commits)"
	}
	out, err := exec.Command("git", "-C", root, "-c", "core.quotePath=false", "status", "--porcelain=v1", "--untracked-files=all").Output()
	if err != nil {
		return GitState{}
	}
	var relevant []string
	for _, line := range strings.Split(strings.TrimRight(string(out), "\n"), "\n") {
		if line == "" {
			continue
		}
		path := line
		if len(path) > 3 {
			path = path[3:]
		}
		if before, after, found := strings.Cut(path, " -> "); found && before != "" {
			path = after
		}
		path = strings.Trim(path, "\"")
		if strings.HasPrefix(strings.Trim(path, "\""), ".soloweave/context/") {
			continue
		}
		relevant = append(relevant, line+":"+workingFileDigest(root, path))
		state.Changes = append(state.Changes, path)
	}
	sum := sha256.Sum256([]byte(strings.Join(relevant, "\n")))
	state.Digest = hex.EncodeToString(sum[:])
	state.SnapshotDigest, err = projectSnapshotDigest(root)
	if err != nil {
		return GitState{}
	}
	return state
}

// projectSnapshotDigest tracks project files instead of commit and staging metadata.
// A commit of unchanged files must not invalidate a handoff checkpoint.
func projectSnapshotDigest(root string) (string, error) {
	out, err := exec.Command("git", "-C", root, "ls-files", "--cached", "--others", "--exclude-standard", "-z", "--").Output()
	if err != nil {
		return "", err
	}
	seen := make(map[string]bool)
	var paths []string
	for _, path := range strings.Split(string(out), "\x00") {
		if path == "" || strings.HasPrefix(path, ".soloweave/context/") || seen[path] {
			continue
		}
		seen[path] = true
		paths = append(paths, path)
	}
	sort.Strings(paths)
	h := sha256.New()
	for _, path := range paths {
		_, _ = io.WriteString(h, path)
		_, _ = h.Write([]byte{0})
		_, _ = io.WriteString(h, workingFileDigest(root, path))
		_, _ = h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func workingFileDigest(root, rel string) string {
	path := filepath.Join(root, filepath.FromSlash(rel))
	info, err := os.Lstat(path)
	if err != nil {
		return "missing"
	}
	if info.Mode()&os.ModeSymlink != 0 {
		link, err := os.Readlink(path)
		if err != nil {
			return "unreadable symlink"
		}
		sum := sha256.Sum256([]byte(link))
		return hex.EncodeToString(sum[:])
	}
	if !info.Mode().IsRegular() {
		return "non-regular"
	}
	file, err := os.Open(path)
	if err != nil {
		return "unreadable"
	}
	defer file.Close()
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "unreadable"
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Checkpoint(root string, input Input) error {
	if strings.TrimSpace(input.Task) == "" || strings.TrimSpace(input.Summary) == "" || strings.TrimSpace(input.Next) == "" {
		return errors.New("task, summary and next step are required")
	}
	if err := Initialize(root, filepath.Base(root)); err != nil {
		return err
	}
	git := Git(root)
	verification := input.Verification
	if len(verification) == 0 {
		verification = []string{"Not run"}
	}
	issues := input.Issues
	if len(issues) == 0 {
		issues = []string{"None recorded"}
	}
	gitText := "Git unavailable"
	if git.Available {
		gitText = fmt.Sprintf("Branch: %s\n\nHEAD: %s\n\nChanged files: %s", git.Branch, git.Head, strings.Join(git.Changes, ", "))
	}
	body := fmt.Sprintf("# Project Handoff\n\n## Current Task\n\n%s\n\n## Current State\n\n%s\n\n## Recent Changes\n\n%s\n\n## Verification\n\n%s\n\n## Known Issues\n\n%s\n\n## Next Steps\n\n%s\n\n## Git State\n\n%s\n", input.Task, input.Summary, strings.Join(git.Changes, ", "), strings.Join(verification, "\n"), strings.Join(issues, "\n"), input.Next, gitText)
	meta, err := json.MarshalIndent(Metadata{At: time.Now().UTC(), Git: git}, "", "  ")
	if err != nil {
		return err
	}
	if err := writeManaged(filepath.Join(dir(root), "HANDOFF.md"), []byte(body), false); err != nil {
		return err
	}
	return writeManaged(filepath.Join(dir(root), "checkpoint.json"), append(meta, '\n'), false)
}

func Check(root string) []string {
	var problems []string
	for _, name := range []string{"PROJECT.md", "STATUS.md", "HANDOFF.md", "CHANGES.md"} {
		if _, err := os.Stat(filepath.Join(dir(root), name)); err != nil {
			problems = append(problems, "missing "+name)
		}
	}
	data, err := os.ReadFile(filepath.Join(dir(root), "checkpoint.json"))
	if errors.Is(err, os.ErrNotExist) {
		return append(problems, "missing checkpoint")
	}
	if err != nil {
		return append(problems, "cannot read checkpoint: "+err.Error())
	}
	var meta Metadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return append(problems, "invalid checkpoint metadata")
	}
	current := Git(root)
	if meta.Git.Available && !current.Available {
		problems = append(problems, "Git unavailable since checkpoint")
	}
	if meta.Git.Available && current.Available {
		stale := meta.Git.Branch != current.Branch
		if meta.Git.SnapshotDigest != "" {
			stale = stale || meta.Git.SnapshotDigest != current.SnapshotDigest
		} else {
			// Checkpoints created before snapshot_digest used commit and status metadata.
			stale = stale || meta.Git.Head != current.Head || meta.Git.Digest != current.Digest
		}
		if stale {
			problems = append(problems, "stale checkpoint: Git state changed")
		}
	}
	return problems
}

func Resume(root string) (string, error) {
	data, err := os.ReadFile(filepath.Join(dir(root), "HANDOFF.md"))
	if err != nil {
		return "", err
	}
	result := string(data)
	if problems := Check(root); len(problems) > 0 {
		result += "\n## Continuity Warnings\n\n" + strings.Join(problems, "\n") + "\n"
	}
	return result, nil
}

func Show(root string) (string, error) {
	var parts []string
	for _, name := range []string{"PROJECT.md", "STATUS.md", "HANDOFF.md"} {
		data, err := os.ReadFile(filepath.Join(dir(root), name))
		if err != nil {
			return "", err
		}
		parts = append(parts, string(data))
	}
	return strings.Join(parts, "\n---\n\n"), nil
}
