package project

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"go.yaml.in/yaml/v3"
	"soloweave/internal/config"
	"soloweave/internal/continuity"
)

type Approval struct {
	Digest string `json:"digest"`
	ADR    string `json:"adr"`
}

func approvalPath(root string) string { return filepath.Join(root, ".soloweave", "approval.json") }

var adrName = regexp.MustCompile(`^ADR-[0-9]{4}\.md$`)

func safeApprovalPath(root string) (string, error) {
	path := approvalPath(root)
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return "", errors.New("approval record is symlink")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	return path, nil
}

func architectureDigest(cfg config.Config) (string, error) {
	value := struct {
		ProjectType string
		Frontend    *config.Stack
		Backend     *config.Stack
		Database    *config.Database
		Workspace   *config.Workspace
	}{cfg.Project.Type, cfg.Frontend, cfg.Backend, cfg.Database, cfg.Workspace}
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func Initialize(root string, cfg config.Config) error {
	if _, err := os.Stat(config.Path(root)); err == nil {
		return errors.New("project already initialized")
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	cfg.Project.Status = "draft"
	if err := config.Validate(cfg); err != nil {
		return err
	}
	if err := continuity.Initialize(root, cfg.Project.Name); err != nil {
		return err
	}
	return config.Save(root, cfg)
}

func readApproval(root string) (Approval, error) {
	path, err := safeApprovalPath(root)
	if err != nil {
		return Approval{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Approval{}, err
	}
	var a Approval
	if err := json.Unmarshal(data, &a); err != nil {
		return Approval{}, err
	}
	return a, nil
}

func CheckApproval(root string) error {
	cfg, err := config.Load(root)
	if err != nil {
		return err
	}
	if cfg.Project.Status != "approved" {
		return errors.New("architecture not approved")
	}
	a, err := readApproval(root)
	if err != nil {
		return fmt.Errorf("approval record unavailable: %w", err)
	}
	digest, err := architectureDigest(cfg)
	if err != nil {
		return err
	}
	if a.Digest != digest {
		return errors.New("approved architecture changed without a new ADR")
	}
	if !adrName.MatchString(a.ADR) {
		return errors.New("invalid approval ADR path")
	}
	if _, err := os.Stat(filepath.Join(root, ".soloweave", "decisions", a.ADR)); err != nil {
		return fmt.Errorf("approval ADR missing: %w", err)
	}
	return nil
}

func Approve(root string) (string, error) {
	approvalFile, err := safeApprovalPath(root)
	if err != nil {
		return "", err
	}
	cfg, err := config.Load(root)
	if err != nil {
		return "", err
	}
	digest, err := architectureDigest(cfg)
	if err != nil {
		return "", err
	}
	if cfg.Project.Status == "approved" {
		if a, err := readApproval(root); err == nil && a.Digest == digest {
			return a.ADR, nil
		}
	}
	dir := filepath.Join(root, ".soloweave", "decisions")
	if info, err := os.Lstat(dir); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return "", errors.New("decisions path is symlink")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	var name string
	for n := 1; ; n++ {
		name = fmt.Sprintf("ADR-%04d.md", n)
		if _, err := os.Stat(filepath.Join(dir, name)); errors.Is(err, os.ErrNotExist) {
			break
		} else if err != nil {
			return "", err
		}
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	body := fmt.Sprintf("# %s: Approved project architecture\n\nDate: %s\n\nStatus: Accepted by local `soloweave approve` command.\n\n## Decision\n\nThe developer approved the following configuration. Major changes need another approval and ADR.\n\n```yaml\n%s```\n", name, time.Now().Format("2006-01-02"), data)
	path := filepath.Join(dir, name)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", err
	}
	if _, err := file.WriteString(body); err != nil {
		file.Close()
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	cfg.Project.Status = "approved"
	if err := config.Save(root, cfg); err != nil {
		return "", err
	}
	approval, err := json.MarshalIndent(Approval{Digest: digest, ADR: name}, "", "  ")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(approvalFile, append(approval, '\n'), 0644); err != nil {
		return "", err
	}
	return name, nil
}
