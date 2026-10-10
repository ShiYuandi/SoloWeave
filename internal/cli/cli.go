package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ShiYuandi/SoloWeave/internal/catalog"
	"github.com/ShiYuandi/SoloWeave/internal/config"
	"github.com/ShiYuandi/SoloWeave/internal/continuity"
	"github.com/ShiYuandi/SoloWeave/internal/installer"
	"github.com/ShiYuandi/SoloWeave/internal/project"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string { return e.Err.Error() }
func (e *ExitError) Unwrap() error { return e.Err }
func ExitCode(err error) int {
	if err == nil {
		return 0
	}
	var e *ExitError
	if errors.As(err, &e) {
		return e.Code
	}
	return 2
}

func New(root string, in io.Reader, out, errOut io.Writer) *cobra.Command {
	reader := bufio.NewReader(in)
	prompt := func(label string) (string, error) {
		fmt.Fprint(out, label)
		value, err := reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}
		return strings.TrimSpace(value), nil
	}
	rootCmd := &cobra.Command{Use: "soloweave", Short: "AI engineering skills and project continuity", SilenceUsage: true, SilenceErrors: true}
	rootCmd.SetIn(in)
	rootCmd.SetOut(out)
	rootCmd.SetErr(errOut)

	var from, preset, name, kind, frontendLanguage, backendLanguage, frontendFramework, backendFramework, agentList string
	var teamSize int
	initCmd := &cobra.Command{Use: "init", Short: "Create a draft project contract", RunE: func(cmd *cobra.Command, args []string) error {
		var cfg config.Config
		if from != "" && preset != "" {
			return errors.New("--from and --preset cannot be combined")
		}
		if from == "" && preset == "" && name == "" && kind == "" {
			items, err := catalog.List()
			if err != nil {
				return err
			}
			fmt.Fprintln(out, "Choose a stack preset or 'custom':")
			for _, item := range items {
				fmt.Fprintf(out, "  %s: %s\n", item.ID, item.Description)
			}
			preset, err = prompt("Preset ID [custom]: ")
			if err != nil {
				return err
			}
			if preset == "custom" {
				preset = ""
			}
		}
		if from != "" {
			data, err := os.ReadFile(from)
			if err != nil {
				return err
			}
			cfg, err = config.Decode(data)
			if err != nil {
				return err
			}
		} else if preset != "" {
			if kind != "" {
				return errors.New("--type cannot be combined with --preset; edit the draft after initialization")
			}
			var err error
			cfg, err = catalog.Preset(preset)
			if err != nil {
				return err
			}
			if name == "" {
				name, err = prompt("Project name: ")
				if err != nil {
					return err
				}
			}
			if name != "" {
				cfg.Project.Name = name
			}
			if cmd.Flags().Changed("team-size") {
				cfg.Project.TeamSize = teamSize
			}
			if agentList != "" {
				cfg.Agents = splitAgents(agentList)
			}
			if cfg.Frontend != nil {
				if frontendLanguage != "" {
					cfg.Frontend.Language = frontendLanguage
				}
				if frontendFramework != "" {
					cfg.Frontend.Framework = frontendFramework
				}
			}
			if cfg.Backend != nil {
				if backendLanguage != "" {
					cfg.Backend.Language = backendLanguage
				}
				if backendFramework != "" {
					cfg.Backend.Framework = backendFramework
				}
			}
		} else {
			var err error
			if name == "" {
				name, err = prompt("Project name: ")
				if err != nil {
					return err
				}
			}
			if kind == "" {
				kind, err = prompt("Project type (frontend/backend/fullstack): ")
				if err != nil {
					return err
				}
			}
			cfg.SchemaVersion = 1
			cfg.Project = config.Project{Name: name, Type: kind, Status: "draft", TeamSize: teamSize}
			if kind == "frontend" || kind == "fullstack" {
				if frontendLanguage == "" {
					frontendLanguage, err = prompt("Frontend language: ")
					if err != nil {
						return err
					}
				}
				cfg.Frontend = &config.Stack{Language: frontendLanguage, Framework: frontendFramework}
			}
			if kind == "backend" || kind == "fullstack" {
				if backendLanguage == "" {
					backendLanguage, err = prompt("Backend language: ")
					if err != nil {
						return err
					}
				}
				cfg.Backend = &config.Stack{Language: backendLanguage, Framework: backendFramework}
			}
			if agentList != "" {
				cfg.Agents = splitAgents(agentList)
			} else {
				cfg.Agents = []string{"codex"}
			}
		}
		if err := project.Initialize(root, cfg); err != nil {
			return err
		}
		fmt.Fprintln(out, "Draft project created at .soloweave/project.yaml. Review it, then run soloweave approve.")
		return nil
	}}
	initCmd.Flags().StringVar(&from, "from", "", "Read a draft YAML configuration")
	initCmd.Flags().StringVar(&preset, "preset", "", "Start from a catalog preset (see catalog)")
	initCmd.Flags().StringVar(&name, "name", "", "Project name")
	initCmd.Flags().StringVar(&kind, "type", "", "frontend, backend, or fullstack")
	initCmd.Flags().StringVar(&frontendLanguage, "frontend-language", "", "Frontend language")
	initCmd.Flags().StringVar(&backendLanguage, "backend-language", "", "Backend language")
	initCmd.Flags().StringVar(&frontendFramework, "frontend-framework", "", "Frontend framework")
	initCmd.Flags().StringVar(&backendFramework, "backend-framework", "", "Backend framework")
	initCmd.Flags().StringVar(&agentList, "agents", "", "Comma-separated agents")
	initCmd.Flags().IntVar(&teamSize, "team-size", 1, "Team size for context")
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(&cobra.Command{Use: "catalog", Short: "List optional stack presets", RunE: func(cmd *cobra.Command, args []string) error {
		items, err := catalog.List()
		if err != nil {
			return err
		}
		for _, item := range items {
			fmt.Fprintf(out, "%s\t%s\n", item.ID, item.Description)
		}
		return nil
	}})

	var yes bool
	approveCmd := &cobra.Command{Use: "approve", Short: "Approve current architecture and record an ADR", RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(root)
		if err != nil {
			return err
		}
		data, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "Architecture to approve:\n%s\n", data)
		if !yes {
			answer, err := prompt("Approve this architecture? [y/N]: ")
			if err != nil {
				return err
			}
			if answer != "y" && answer != "yes" {
				return &ExitError{Code: 2, Err: errors.New("approval declined")}
			}
		}
		adr, err := project.Approve(root)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, "Approved:", adr)
		return nil
	}}
	approveCmd.Flags().BoolVar(&yes, "yes", false, "Explicitly confirm architecture without prompt")
	rootCmd.AddCommand(approveCmd)

	var selected string
	var dryRun bool
	installCmd := &cobra.Command{Use: "install", Short: "Install skills and platform rules", RunE: func(cmd *cobra.Command, args []string) error {
		if selected == "" {
			cfg, err := config.Load(root)
			if err != nil {
				return err
			}
			selected = strings.Join(cfg.Agents, ",")
		}
		changes, err := installer.Install(root, splitAgents(selected), dryRun)
		if err != nil {
			return err
		}
		for _, change := range changes {
			fmt.Fprintf(out, "%s %s\n", strings.ToUpper(change.Action), change.Path)
		}
		if len(changes) == 0 {
			fmt.Fprintln(out, "Already installed; no changes.")
		}
		if dryRun {
			fmt.Fprintln(out, "Dry run: no files written.")
		}
		return nil
	}}
	installCmd.Flags().StringVar(&selected, "agents", "", "Comma-separated agents (default: project config)")
	installCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without writing")
	rootCmd.AddCommand(installCmd)

	rootCmd.AddCommand(&cobra.Command{Use: "version", Short: "Print version", Run: func(cmd *cobra.Command, args []string) { fmt.Fprintln(out, installer.Version) }})
	rootCmd.AddCommand(&cobra.Command{Use: "check", Short: "Check the project contract and handoff", RunE: func(cmd *cobra.Command, args []string) error {
		failed := false
		if _, err := config.Load(root); err != nil {
			fmt.Fprintln(out, "FAIL config:", err)
			failed = true
		} else {
			fmt.Fprintln(out, "PASS config")
		}
		if err := project.CheckApproval(root); err != nil {
			fmt.Fprintln(out, "FAIL approval:", err)
			failed = true
		} else {
			fmt.Fprintln(out, "PASS approval")
		}
		if problems := installer.Check(root); len(problems) > 0 {
			for _, p := range problems {
				fmt.Fprintln(out, "FAIL install:", p)
			}
			failed = true
		} else {
			fmt.Fprintln(out, "PASS install")
		}
		if problems := continuity.Check(root); len(problems) > 0 {
			for _, p := range problems {
				fmt.Fprintln(out, "FAIL context:", p)
			}
			failed = true
		} else {
			fmt.Fprintln(out, "PASS context")
		}
		if !continuity.Git(root).Available {
			fmt.Fprintln(out, "UNAVAILABLE Git repository: Git state not verified")
		}
		if failed {
			return &ExitError{Code: 1, Err: errors.New("project check failed")}
		}
		return nil
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "doctor", Short: "Report setup and installation health", RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := exec.LookPath("git"); err != nil {
			fmt.Fprintln(out, "UNAVAILABLE git executable")
		} else {
			fmt.Fprintln(out, "PASS git executable")
		}
		if !continuity.Git(root).Available {
			fmt.Fprintln(out, "UNAVAILABLE Git repository")
		} else {
			fmt.Fprintln(out, "PASS Git repository")
		}
		if _, err := config.Load(root); err != nil {
			fmt.Fprintln(out, "NOT RUN project checks: no valid project contract")
			fmt.Fprintln(out, "NOT RUN install: no valid project contract")
			return nil
		}
		fmt.Fprintln(out, "PASS project config")
		problems := installer.Check(root)
		for _, p := range problems {
			fmt.Fprintln(out, "FAIL install:", p)
		}
		if len(problems) > 0 {
			return &ExitError{Code: 1, Err: errors.New("installation check failed")}
		}
		fmt.Fprintln(out, "PASS install")
		return nil
	}})

	contextCmd := &cobra.Command{Use: "context", Short: "Inspect and update project continuity"}
	contextCmd.AddCommand(&cobra.Command{Use: "show", RunE: func(cmd *cobra.Command, args []string) error {
		text, err := continuity.Show(root)
		if err == nil {
			fmt.Fprintln(out, text)
		}
		return err
	}})
	contextCmd.AddCommand(&cobra.Command{Use: "resume", RunE: func(cmd *cobra.Command, args []string) error {
		text, err := continuity.Resume(root)
		if err == nil {
			fmt.Fprintln(out, text)
		}
		return err
	}})
	contextCmd.AddCommand(&cobra.Command{Use: "check", RunE: func(cmd *cobra.Command, args []string) error {
		problems := continuity.Check(root)
		for _, p := range problems {
			fmt.Fprintln(out, "FAIL", p)
		}
		if !continuity.Git(root).Available {
			fmt.Fprintln(out, "UNAVAILABLE Git repository")
		}
		if len(problems) > 0 {
			return &ExitError{Code: 1, Err: errors.New("context check failed")}
		}
		fmt.Fprintln(out, "PASS context documents")
		return nil
	}})
	var task, summary, next string
	var verification, issues []string
	checkpointCmd := &cobra.Command{Use: "checkpoint", RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if task == "" {
			task, err = prompt("Current task: ")
			if err != nil {
				return err
			}
		}
		if summary == "" {
			summary, err = prompt("Current state / recent changes: ")
			if err != nil {
				return err
			}
		}
		if next == "" {
			next, err = prompt("Next step: ")
			if err != nil {
				return err
			}
		}
		if err := continuity.Checkpoint(root, continuity.Input{Task: task, Summary: summary, Next: next, Verification: verification, Issues: issues}); err != nil {
			return err
		}
		fmt.Fprintln(out, "Checkpoint saved.")
		if !continuity.Git(root).Available {
			fmt.Fprintln(out, "UNAVAILABLE Git repository: checkpoint has no Git evidence")
		}
		return nil
	}}
	checkpointCmd.Flags().StringVar(&task, "task", "", "Current task")
	checkpointCmd.Flags().StringVar(&summary, "summary", "", "Completed work and current state")
	checkpointCmd.Flags().StringVar(&next, "next", "", "Next action")
	checkpointCmd.Flags().StringSliceVar(&verification, "verification", nil, "Actual check and result; repeatable")
	checkpointCmd.Flags().StringSliceVar(&issues, "issue", nil, "Known issue; repeatable")
	contextCmd.AddCommand(checkpointCmd)
	rootCmd.AddCommand(contextCmd)
	return rootCmd
}

func splitAgents(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
