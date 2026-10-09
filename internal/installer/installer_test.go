package installer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"soloweave/internal/config"
	"soloweave/internal/project"
)

func approvedProject(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	cfg := config.Config{SchemaVersion: 1}
	cfg.Project = config.Project{Name: "demo", Type: "backend", Status: "draft"}
	cfg.Backend = &config.Stack{Language: "go"}
	cfg.Agents = []string{"codex", "claude", "cursor"}
	if err := project.Initialize(root, cfg); err != nil {
		t.Fatal(err)
	}
	if _, err := project.Approve(root); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestDryRunInstallAndRepeat(t *testing.T) {
	root := approvedProject(t)
	changes, err := Install(root, []string{"codex", "claude", "cursor"}, true)
	if err != nil || len(changes) == 0 {
		t.Fatalf("dry run: %v %v", changes, err)
	}
	if _, err := os.Stat(filepath.Join(root, "AGENTS.md")); !os.IsNotExist(err) {
		t.Fatalf("dry run wrote rules: %v", err)
	}
	if _, err := Install(root, []string{"codex", "claude", "cursor"}, false); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{".agents/skills/soloweave/SKILL.md", ".claude/skills/project-continuity/SKILL.md", ".cursor/rules/soloweave.mdc", ".soloweave/installation.json"} {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(path))); err != nil {
			t.Fatalf("missing %s: %v", path, err)
		}
	}
	changes, err = Install(root, []string{"codex", "claude", "cursor"}, false)
	if err != nil || len(changes) != 0 {
		t.Fatalf("repeat changed files: %v %v", changes, err)
	}
}

func TestExistingRulesPreservedAndModifiedSkillConflicts(t *testing.T) {
	root := approvedProject(t)
	rules := filepath.Join(root, "AGENTS.md")
	if err := os.WriteFile(rules, []byte("User rules\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := Install(root, []string{"codex"}, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(rules)
	if err != nil || !strings.HasPrefix(string(data), "User rules\n") {
		t.Fatalf("user rules lost: %s, %v", data, err)
	}
	skill := filepath.Join(root, ".agents", "skills", "soloweave", "SKILL.md")
	if err := os.WriteFile(skill, []byte("user edit"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := Install(root, []string{"codex"}, false); err == nil {
		t.Fatal("modified managed skill should conflict")
	}
	data, _ = os.ReadFile(skill)
	if string(data) != "user edit" {
		t.Fatal("conflict overwrote user edit")
	}
}

func TestAddingAgentPreservesPreviouslyInstalledAgent(t *testing.T) {
	root := approvedProject(t)
	if _, err := Install(root, []string{"codex"}, false); err != nil {
		t.Fatal(err)
	}
	if _, err := Install(root, []string{"claude"}, false); err != nil {
		t.Fatal(err)
	}
	m, err := readManifest(root)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(m.Agents, ",") != "claude,codex" {
		t.Fatalf("agents lost: %v", m.Agents)
	}
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("installation unhealthy: %v", problems)
	}
}

func TestCheckDetectsIncompleteManifest(t *testing.T) {
	root := approvedProject(t)
	if _, err := Install(root, []string{"codex"}, false); err != nil {
		t.Fatal(err)
	}
	m, err := readManifest(root)
	if err != nil {
		t.Fatal(err)
	}
	delete(m.Files, ".agents/skills/soloweave/SKILL.md")
	data, _ := json.Marshal(m)
	if err := os.WriteFile(manifestPath(root), data, 0644); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) == 0 {
		t.Fatal("incomplete manifest accepted")
	}
}
