# Repository layout

[简体中文](../directory-structure.md)

SoloWeave provides AI development Skills.

| Path | Purpose |
| --- | --- |
| `skills/` | The sole source for five installable Skills. |
| `skills/project-setup/references/` | On-demand selection guide installed with the Skill. |
| `scripts/check-skills.ps1`, `scripts/check-docs.ps1` | Metadata, resources, nonempty Markdown, and relative file link checks. |
| `scripts/test-tooling.ps1` | Negative validation cases and package regression checks. |
| `scripts/prepare-behavior-tests.ps1`, `scripts/run-codex-behavior-tests.ps1` | Isolated fixtures and real Codex sessions; acceptance requires review. |
| `scripts/package-skillhub.ps1` | Reproducible five ZIPs and SHA256SUMS. |
| `.github/workflows/ci.yml` | Runs maintenance checks and uploads candidate packages. |
| `README.md`, `CONTRIBUTING.md`, `docs/en/` | Project entry, contribution rules, and English public documentation. |
| `docs/design.md`, `docs/development-plan.md` | Current product design and development plan. |
| `docs/configuration.md`, `docs/project-continuity.md` | Project records, decisions, and handoff. |
| `docs/agent-acceptance.md`, `docs/behavior-testing.md`, `docs/five-skills-acceptance.md`, `docs/end-to-end-acceptance.md` | Acceptance criteria, live-test method, and current five-Skill results; historical evidence remains separate. |
| `docs/codex-evidence/`, `docs/claude-code-evidence/` | Reviewed receipts or original client evidence. |
| `docs/archive/` | Historical selection design and completed implementation plans; current rules are in docs/design.md and skills/. |
| `docs/releasing.md`, `docs/release-status.md` | Publication workflow and verified online status. |
| `.soloweave/context/`, `.soloweave/decisions/` | Repository continuity and architecture decisions. |
| `AGENTS.md` | Repository rules for agents. |
| `LICENSE`, `SECURITY.md`, `CODE_OF_CONDUCT.md` | Licensing, security, and community policy. |
| `dist/skills/`, `dist/acceptance/` | Ignored current packages and original local test evidence. |
| `dist/archive/` | Recoverable old package snapshots, fixtures, and local tool remnants; excluded from Git. |
| `.tools/` | Optional ignored, regenerable tool cache; removed during cleanup. |
| `.git/` | Git metadata. |

Target projects may reuse existing documentation or store context under `.soloweave/context/` and `.soloweave/decisions/`. Check existing installations and preserve user edits before installing or updating.
