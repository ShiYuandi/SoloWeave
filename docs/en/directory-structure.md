# Repository directory guide

[简体中文](../directory-structure.md)

This page describes the **SoloWeave source repository**. For directories created in a target project, see [configuration](configuration.md) and [project continuity](project-continuity.md).

## Repository root

| Path | Purpose | Commit? |
| --- | --- | --- |
| `README.md` | Chinese project introduction, download, and quick start | Yes |
| `AGENTS.md` | Repository guidance for AI coding agents | Yes |
| `go.mod`, `go.sum` | Go module declaration and dependency checksums | Yes |
| `LICENSE` | Original MIT license text | Yes |
| `CONTRIBUTING.md` | Contribution and verification guidance | Yes |
| `SECURITY.md` | Security reporting and secret-handling guidance | Yes |
| `CODE_OF_CONDUCT.md` | Community conduct | Yes |
| `CHANGELOG.md` | Version history | Yes |
| `.gitignore` | Excludes local tools, build outputs, and temporary files | Yes |
| `.gitattributes` | Normalizes text line endings across platforms | Yes |
| `.git/` | Local Git metadata, managed by the repository owner | No |
| `.tools/` | Local Go toolchain, caches, validators, and other reproducible tools | No |
| `dist/` | Windows download package, local preview binaries, and `SHA256SUMS` | No |

`.tools/` and `dist/` stay on the development machine for continued work and verification. Cross-compiled artifacts in `dist/` do not prove that the binaries run on every target OS.

## Source and build

| Path | Purpose |
| --- | --- |
| `cmd/soloweave/main.go` | CLI entry point and exit-code handling |
| `internal/cli/` | Cobra commands, prompts, and end-to-end CLI tests |
| `internal/config/` | YAML read/write, JSON Schema and Go semantic validation; `schema/project.schema.json` defines the structure |
| `internal/project/` | Initialization, architecture approval, approval digest, and ADRs |
| `internal/continuity/` | Context documents, Git state, checkpoints, and recovery |
| `internal/catalog/` | Optional stack preset catalog |
| `internal/installer/` | Installation and checking for Codex, Claude Code, and Cursor, including preview and conflict handling |
| `internal/bundle/` | Go-embedded resources: presets in `assets/catalogs/`, platform rules in `assets/rules/`, and the single source of five Skills in `assets/skills/` |
| `scripts/build-preview.ps1` | Builds four preview targets and checksum file |
| `scripts/package-windows.ps1` | Creates a Windows x64 ZIP with bilingual guides and a SHA256 checksum |
| `.github/workflows/ci.yml` | Go tests, vet, and build on a remote repository |
| `.github/workflows/release-windows.yml` | Manually tests and packages Windows, then creates a draft Release |

Each `*_test.go` file tests its neighboring package. Rebuild the CLI after changing embedded assets; an older binary still contains its previous copies.

## Design and handoff

| Path | Purpose |
| --- | --- |
| `docs/development-plan.md` | V0.1 scope, steps, and acceptance record; available in Chinese |
| `docs/design-v0.2.md` | Product design baseline approved by the user; available in Chinese |
| `docs/configuration.md` | Configuration fields and approval flow |
| `docs/project-continuity.md` | Handoff files and checkpoints |
| `docs/skills.md` | Five Skills and platform installation paths |
| `docs/skill-only-workflow.md` | Planning, checks, and handoffs with Skills only and no CLI |
| `docs/package-readme.zh-CN.md`, `docs/package-readme.en.md` | Versioned Chinese and English README templates included in the Windows ZIP |
| `docs/agent-acceptance.md` | Pending live client handoff acceptance steps |
| `docs/directory-structure.md` | Chinese version of this guide |
| `docs/en/` | English versions of public reader-facing documents |
| `.soloweave/project.yaml` | This repository's own approved project contract |
| `.soloweave/approval.json` | CLI-managed approval digest |
| `.soloweave/decisions/ADR-0001.md` | This repository's architecture decision |
| `.soloweave/context/PROJECT.md` | Stable project goals and entry points |
| `.soloweave/context/STATUS.md` | Completed, active, blocked, and planned work |
| `.soloweave/context/HANDOFF.md` | Current task, actual checks, risks, and next step |
| `.soloweave/context/CHANGES.md` | Important changes |
| `.soloweave/context/checkpoint.json` | CLI-managed machine-readable Git checkpoint data |
| `.soloweave/context/tasks/` | Optional durable task notes; an empty directory is not committed |

`.soloweave/` is part of the project's durable knowledge and should be retained with source files. A checkpoint records the Git branch and project file contents. Editing a project file makes `context check` report staleness; committing identical contents does not invalidate a new checkpoint.
