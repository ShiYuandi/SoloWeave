package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func run(t *testing.T, root, input string, args ...string) (string, error) {
	t.Helper()
	var output bytes.Buffer
	cmd := New(root, strings.NewReader(input), &output, &output)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return output.String(), err
}

func TestEndToEndLocalProject(t *testing.T) {
	root := t.TempDir()
	seed := filepath.Join(root, "seed.yaml")
	yaml := "schema_version: 1\nproject: {name: demo, type: backend, status: draft, team_size: 1}\nbackend: {language: go, framework: custom}\nagents: [codex, claude, cursor]\n"
	if err := os.WriteFile(seed, []byte(yaml), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "", "init", "--from", seed); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "n\n", "approve"); err == nil {
		t.Fatal("declined approval succeeded")
	}
	if _, err := run(t, root, "", "approve", "--yes"); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "", "install", "--agents", "codex,claude,cursor", "--dry-run"); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "", "install", "--agents", "codex,claude,cursor"); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "", "context", "checkpoint", "--task", "login", "--summary", "form complete", "--next", "wire API"); err != nil {
		t.Fatal(err)
	}
	output, err := run(t, root, "", "context", "resume")
	if err != nil || !strings.Contains(output, "wire API") || !strings.Contains(output, "Git unavailable") {
		t.Fatalf("resume: %q %v", output, err)
	}
	output, err = run(t, root, "", "check")
	if err != nil || !strings.Contains(output, "PASS") {
		t.Fatalf("check: %q %v", output, err)
	}
}

func TestInitPresetCanBeRenamedAndCatalogListed(t *testing.T) {
	root := t.TempDir()
	output, err := run(t, root, "", "catalog")
	if err != nil || !strings.Contains(output, "api-go") {
		t.Fatalf("catalog: %q %v", output, err)
	}
	if _, err := run(t, root, "", "init", "--preset", "api-go", "--name", "my-api"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(root, ".soloweave", "project.yaml"))
	if err != nil || !strings.Contains(string(data), "my-api") || !strings.Contains(string(data), "status: draft") {
		t.Fatalf("preset config: %s %v", data, err)
	}
}

func TestInteractiveInitOffersPreset(t *testing.T) {
	root := t.TempDir()
	if _, err := run(t, root, "api-go\ninteractive-api\n", "init"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(root, ".soloweave", "project.yaml"))
	if err != nil || !strings.Contains(string(data), "interactive-api") || !strings.Contains(string(data), "language: go") {
		t.Fatalf("interactive preset: %s %v", data, err)
	}
}

func TestDoctorDistinguishesNotRunAndFailedInstallation(t *testing.T) {
	root := t.TempDir()
	output, err := run(t, root, "", "doctor")
	if err != nil || !strings.Contains(output, "NOT RUN project checks") || !strings.Contains(output, "NOT RUN install") {
		t.Fatalf("uninitialized doctor: %q %v", output, err)
	}
	if _, err := run(t, root, "", "init", "--preset", "api-go", "--name", "demo"); err != nil {
		t.Fatal(err)
	}
	if _, err := run(t, root, "", "approve", "--yes"); err != nil {
		t.Fatal(err)
	}
	output, err = run(t, root, "", "doctor")
	if err == nil || !strings.Contains(output, "FAIL install") {
		t.Fatalf("missing install: %q %v", output, err)
	}
}
