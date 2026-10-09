package config

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.yaml.in/yaml/v3"
)

//go:embed schema/project.schema.json
var schemaFS embed.FS

type Config struct {
	SchemaVersion int             `yaml:"schema_version" json:"schema_version"`
	Project       Project         `yaml:"project" json:"project"`
	Frontend      *Stack          `yaml:"frontend,omitempty" json:"frontend,omitempty"`
	Backend       *Stack          `yaml:"backend,omitempty" json:"backend,omitempty"`
	Database      *Database       `yaml:"database,omitempty" json:"database,omitempty"`
	Workspace     *Workspace      `yaml:"workspace,omitempty" json:"workspace,omitempty"`
	Reuse         map[string]bool `yaml:"reuse,omitempty" json:"reuse,omitempty"`
	Continuity    map[string]bool `yaml:"continuity,omitempty" json:"continuity,omitempty"`
	Quality       *Quality        `yaml:"quality,omitempty" json:"quality,omitempty"`
	Agents        []string        `yaml:"agents,omitempty" json:"agents,omitempty"`
}

type Project struct {
	Name     string `yaml:"name" json:"name"`
	Type     string `yaml:"type" json:"type"`
	Status   string `yaml:"status" json:"status"`
	TeamSize int    `yaml:"team_size,omitempty" json:"team_size,omitempty"`
}

type Stack struct {
	Language       string `yaml:"language" json:"language"`
	Framework      string `yaml:"framework,omitempty" json:"framework,omitempty"`
	PackageManager string `yaml:"package_manager,omitempty" json:"package_manager,omitempty"`
	Architecture   string `yaml:"architecture,omitempty" json:"architecture,omitempty"`
}

type Database struct {
	Engine string `yaml:"engine,omitempty" json:"engine,omitempty"`
	ORM    string `yaml:"orm,omitempty" json:"orm,omitempty"`
}

type Workspace struct {
	Layout       string `yaml:"layout,omitempty" json:"layout,omitempty"`
	FrontendPath string `yaml:"frontend_path,omitempty" json:"frontend_path,omitempty"`
	BackendPath  string `yaml:"backend_path,omitempty" json:"backend_path,omitempty"`
	SharedPath   string `yaml:"shared_path,omitempty" json:"shared_path,omitempty"`
}

type Quality struct {
	Lint          string `yaml:"lint,omitempty" json:"lint,omitempty"`
	RelevantTests string `yaml:"relevant_tests,omitempty" json:"relevant_tests,omitempty"`
}

func Path(root string) string { return filepath.Join(root, ".soloweave", "project.yaml") }

func Decode(data []byte) (Config, error) {
	var raw map[string]any
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return Config{}, fmt.Errorf("invalid YAML: %w", err)
	}
	value, err := json.Marshal(raw)
	if err != nil {
		return Config{}, fmt.Errorf("invalid YAML value: %w", err)
	}
	var doc any
	if err := json.Unmarshal(value, &doc); err != nil {
		return Config{}, err
	}
	if err := validateSchema(doc); err != nil {
		return Config{}, fmt.Errorf("invalid project config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("invalid project config: %w", err)
	}
	if err := Validate(cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func validateSchema(doc any) error {
	data, err := schemaFS.ReadFile("schema/project.schema.json")
	if err != nil {
		return err
	}
	var definition any
	if err := json.Unmarshal(data, &definition); err != nil {
		return err
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource("project.schema.json", definition); err != nil {
		return err
	}
	sch, err := c.Compile("project.schema.json")
	if err != nil {
		return err
	}
	return sch.Validate(doc)
}

func Validate(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	var doc any
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	if err := validateSchema(doc); err != nil {
		return fmt.Errorf("invalid project config: %w", err)
	}
	switch cfg.Project.Type {
	case "frontend":
		if cfg.Frontend == nil || cfg.Backend != nil {
			return errors.New("invalid project config: frontend project requires frontend and no backend")
		}
	case "backend":
		if cfg.Backend == nil || cfg.Frontend != nil {
			return errors.New("invalid project config: backend project requires backend and no frontend")
		}
	case "fullstack":
		if cfg.Frontend == nil || cfg.Backend == nil {
			return errors.New("invalid project config: fullstack project requires frontend and backend")
		}
	}
	if cfg.Frontend != nil && strings.EqualFold(cfg.Frontend.Framework, "nextjs") && !isJS(cfg.Frontend.Language) {
		return errors.New("invalid project config: nextjs requires JavaScript or TypeScript frontend")
	}
	if cfg.Backend != nil && strings.EqualFold(cfg.Backend.Framework, "nestjs") && !strings.EqualFold(cfg.Backend.Language, "typescript") {
		return errors.New("invalid project config: nestjs requires TypeScript backend")
	}
	if cfg.Database != nil && strings.EqualFold(cfg.Database.ORM, "prisma") && (cfg.Backend == nil || !isJS(cfg.Backend.Language)) {
		return errors.New("invalid project config: prisma requires JavaScript or TypeScript backend")
	}
	if cfg.Workspace != nil {
		for _, p := range []string{cfg.Workspace.FrontendPath, cfg.Workspace.BackendPath, cfg.Workspace.SharedPath} {
			if p == "" {
				continue
			}
			clean := filepath.Clean(p)
			if filepath.IsAbs(p) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) || strings.Contains(p, ":") {
				return fmt.Errorf("invalid project config: workspace path %q escapes project", p)
			}
		}
	}
	return nil
}

func isJS(value string) bool {
	return strings.EqualFold(value, "javascript") || strings.EqualFold(value, "typescript")
}

func Load(root string) (Config, error) {
	data, err := os.ReadFile(Path(root))
	if err != nil {
		return Config{}, err
	}
	return Decode(data)
}

func Save(root string, cfg Config) error {
	if err := Validate(cfg); err != nil {
		return err
	}
	dir := filepath.Join(root, ".soloweave")
	if info, err := os.Lstat(dir); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return errors.New("invalid project config directory: symlink")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	path := Path(root)
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return errors.New("invalid project config file: symlink")
	}
	return os.WriteFile(path, bytes.TrimSpace(data), 0644)
}
