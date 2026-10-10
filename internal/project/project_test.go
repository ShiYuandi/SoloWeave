package project

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ShiYuandi/SoloWeave/internal/config"
)

func TestInitializeApproveAndDetectArchitectureDrift(t *testing.T) {
	root := t.TempDir()
	cfg := config.Config{SchemaVersion: 1}
	cfg.Project = config.Project{Name: "demo", Type: "backend", Status: "draft", TeamSize: 1}
	cfg.Backend = &config.Stack{Language: "go", Framework: "custom"}
	if err := Initialize(root, cfg); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".soloweave", "context", "PROJECT.md")); err != nil {
		t.Fatal(err)
	}
	if _, err := Approve(root); err != nil {
		t.Fatal(err)
	}
	if err := CheckApproval(root); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".soloweave", "decisions", "ADR-0001.md")); err != nil {
		t.Fatal(err)
	}
	cfg, _ = config.Load(root)
	cfg.Backend.Framework = "changed"
	if err := config.Save(root, cfg); err != nil {
		t.Fatal(err)
	}
	if err := CheckApproval(root); err == nil {
		t.Fatal("changed architecture accepted")
	}
	if _, err := Approve(root); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, ".soloweave", "decisions", "ADR-0002.md")); err != nil {
		t.Fatal(err)
	}
	if err := CheckApproval(root); err != nil {
		t.Fatal(err)
	}
}

func TestApprovalRejectsADRPathTraversal(t *testing.T) {
	root := t.TempDir()
	cfg := config.Config{SchemaVersion: 1, Project: config.Project{Name: "demo", Type: "backend", Status: "draft"}, Backend: &config.Stack{Language: "go"}}
	if err := Initialize(root, cfg); err != nil {
		t.Fatal(err)
	}
	if _, err := Approve(root); err != nil {
		t.Fatal(err)
	}
	a, err := readApproval(root)
	if err != nil {
		t.Fatal(err)
	}
	a.ADR = "../../outside.md"
	data, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(approvalPath(root), data, 0644); err != nil {
		t.Fatal(err)
	}
	if err := CheckApproval(root); err == nil {
		t.Fatal("approval accepted an ADR outside the project")
	}
}
