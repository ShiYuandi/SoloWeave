package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigRoundTripAndValidation(t *testing.T) {
	root := t.TempDir()
	cfg := Config{SchemaVersion: 1}
	cfg.Project.Name = "demo"
	cfg.Project.Type = "frontend"
	cfg.Project.Status = "draft"
	cfg.Project.TeamSize = 1
	cfg.Frontend = &Stack{Language: "typescript", Framework: "custom-ui"}
	cfg.Agents = []string{"codex"}
	if err := Save(root, cfg); err != nil {
		t.Fatal(err)
	}
	got, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if got.Project.Name != "demo" || got.Frontend.Framework != "custom-ui" {
		t.Fatalf("round trip: %#v", got)
	}
	if _, err := os.Stat(filepath.Join(root, ".soloweave", "project.yaml")); err != nil {
		t.Fatal(err)
	}
}

func TestConfigRejectsInvalidShapeAndCombination(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, ".soloweave", "project.yaml")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	for _, input := range []string{
		"schema_version: 2\nproject: {name: demo, type: frontend, status: draft}\nfrontend: {language: typescript}\n",
		"schema_version: 1\nproject: {name: demo, type: frontend, status: draft}\nbackend: {language: go}\n",
		"schema_version: 1\nproject: {name: demo, type: frontend, status: draft}\nfrontend: {language: typescript}\nworkspace: {frontend_path: ../outside}\n",
	} {
		if err := os.WriteFile(path, []byte(input), 0644); err != nil {
			t.Fatal(err)
		}
		if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "invalid") {
			t.Fatalf("wanted validation error for %q, got %v", input, err)
		}
	}
}

func TestKnownFrameworkCompatibilityWithoutBlockingCustomStacks(t *testing.T) {
	cfg := Config{SchemaVersion: 1}
	cfg.Project = Project{Name: "demo", Type: "backend", Status: "draft"}
	cfg.Backend = &Stack{Language: "go", Framework: "nestjs"}
	if err := Validate(cfg); err == nil {
		t.Fatal("NestJS with Go was accepted")
	}
	cfg.Backend.Framework = "my-go-framework"
	if err := Validate(cfg); err != nil {
		t.Fatalf("custom framework rejected: %v", err)
	}
}
