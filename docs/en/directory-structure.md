# Repository layout

[简体中文](../directory-structure.md)

SoloWeave is a Skills-only project. The current source no longer contains the Go CLI.

| Path | Purpose |
| --- | --- |
| `skills/` | The sole source for six installable Skills. |
| `skills/project-setup/references/` | On-demand selection guide, installed with project-setup. |
| `scripts/check-skills.ps1` | Metadata, resource, and boundary validation. |
| `scripts/check-docs.ps1` | Relative file link validation in Markdown. |
| `scripts/test-tooling.ps1` | Negative validation cases and package regression checks. |
| `scripts/prepare-behavior-tests.ps1`, `scripts/run-codex-behavior-tests.ps1` | Isolated fixtures and real Codex sessions; acceptance requires review. |
| `scripts/package-skillhub.ps1` | Reproducible six ZIPs and SHA256SUMS. |
| `.github/workflows/ci.yml` | Runs the same checks and uploads candidate packages. |
| `docs/en/` | English public documentation. |
| `docs/configuration.md`, `docs/project-continuity.md` | Project records, decisions, and handoff. |
| `docs/cli-migration.md`, `docs/releasing.md` | Legacy migration and release instructions. |
| `docs/archive/go-cli/` | Historical CLI guides and obsolete repository machine records. |
| `docs/design-v0.2.md`, `docs/development-plan.md` | Early CLI design and development history. |
| `.soloweave/context/`, `.soloweave/decisions/` | Repository continuity and ADRs, including Go retirement in ADR-0003. |
| `AGENTS.md` | Repository rules for agents. |
| `LICENSE`, `SECURITY.md`, `CODE_OF_CONDUCT.md` | Licensing, security, and community policy. |
| `.tools/`, `dist/` | Ignored local caches, packages, and test evidence. |
| `.git/` | Owner-managed Git metadata. |

There is no current `cmd/`, `internal/`, `go.mod`, or Go build workflow. Historical commits and Release tags retain the old source. Target projects may reuse existing documentation or use Markdown under `.soloweave/context/` and `.soloweave/decisions/`.
