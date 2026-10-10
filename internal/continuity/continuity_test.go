package continuity

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckpointWithoutGitIsHonestAndRecoverable(t *testing.T) {
	root := t.TempDir()
	if err := Initialize(root, "demo"); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "login", Summary: "form done", Next: "connect API"}); err != nil {
		t.Fatal(err)
	}
	text, err := Resume(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"login", "form done", "connect API", "Git unavailable", "Not run"} {
		if !strings.Contains(text, want) {
			t.Fatalf("resume missing %q:\n%s", want, text)
		}
	}
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("unexpected problems: %v", problems)
	}
}

func TestCheckpointDetectsChangedGitWorkingTree(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	if err := os.WriteFile(filepath.Join(root, "app.txt"), []byte("before"), 0644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "app.txt")
	runGit(t, root, "commit", "-m", "initial")
	if err := Initialize(root, "demo"); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "login", Summary: "done", Next: "review", Verification: []string{"go test ./...: passed"}}); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("fresh checkpoint reported stale: %v", problems)
	}
	if err := os.WriteFile(filepath.Join(root, "app.txt"), []byte("after"), 0644); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) == 0 || !strings.Contains(strings.Join(problems, " "), "stale") {
		t.Fatalf("wanted stale: %v", problems)
	}
}

func TestDirtyCheckpointDetectsFurtherEditToSameFile(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	path := filepath.Join(root, "app.txt")
	if err := os.WriteFile(path, []byte("initial"), 0644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "app.txt")
	runGit(t, root, "commit", "-m", "initial")
	if err := os.WriteFile(path, []byte("dirty once"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := Initialize(root, "demo"); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "feature", Summary: "partial", Next: "continue"}); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("fresh dirty checkpoint: %v", problems)
	}
	if err := os.WriteFile(path, []byte("dirty twice"), 0644); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) == 0 {
		t.Fatal("further edit to dirty file was not detected")
	}
}

func TestCheckpointInUnbornRepositoryReportsNoCommit(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	if err := Initialize(root, "demo"); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "start", Summary: "docs", Next: "code"}); err != nil {
		t.Fatal(err)
	}
	text, err := Resume(root)
	if err != nil || !strings.Contains(text, "no commits") {
		t.Fatalf("unborn repository: %q %v", text, err)
	}
}

func TestCheckpointRemainsFreshAfterCommittingSameFiles(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	if err := os.WriteFile(filepath.Join(root, "app.txt"), []byte("finished"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "feature", Summary: "finished", Next: "review"}); err != nil {
		t.Fatal(err)
	}
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("fresh checkpoint: %v", problems)
	}
	runGit(t, root, "add", "--all")
	runGit(t, root, "commit", "-m", "save finished work")
	if problems := Check(root); len(problems) != 0 {
		t.Fatalf("committing identical files made checkpoint stale: %v", problems)
	}
}

func TestCheckpointDetectsChangedCommittedFiles(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	path := filepath.Join(root, "app.txt")
	if err := os.WriteFile(path, []byte("before"), 0644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "app.txt")
	runGit(t, root, "commit", "-m", "initial")
	if err := Checkpoint(root, Input{Task: "feature", Summary: "done", Next: "review"}); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("after"), 0644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "app.txt")
	runGit(t, root, "commit", "-m", "change content")
	if problems := Check(root); len(problems) == 0 || !strings.Contains(strings.Join(problems, " "), "stale") {
		t.Fatalf("committed content change not detected: %v", problems)
	}
}

func TestLegacyCheckpointStillUsesGitMetadata(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	if err := os.WriteFile(filepath.Join(root, "app.txt"), []byte("finished"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := Checkpoint(root, Input{Task: "feature", Summary: "done", Next: "review"}); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir(root), "checkpoint.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var meta Metadata
	if err := json.Unmarshal(data, &meta); err != nil {
		t.Fatal(err)
	}
	meta.Git.SnapshotDigest = ""
	data, err = json.Marshal(meta)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "--all")
	runGit(t, root, "commit", "-m", "save finished work")
	if problems := Check(root); len(problems) == 0 || !strings.Contains(strings.Join(problems, " "), "stale") {
		t.Fatalf("legacy checkpoint did not retain its Git metadata check: %v", problems)
	}
}

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
}
