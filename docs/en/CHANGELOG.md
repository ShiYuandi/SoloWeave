# Changelog

[简体中文](../../CHANGELOG.md)

## Unreleased: Skills-first migration

- Retired the Go CLI source, dependencies, embedded legacy Skills, builds, and executable release workflow. Archived obsolete machine records and guides; added bilingual migration and Skills release instructions.
- Added PowerShell 7 validation, documentation link checks, package regression tests, and Skills-only CI. Fixed ZIP timestamps make repeated packaging reproducible.
- Removed CLI coupling: entry/feature `0.2.3`, setup `0.3.2`, continuity `0.2.2`, review `0.2.1`; debugging remains `0.1.0`. Independent installation file checks passed; independent live behavior remains unverified.
- Imported Claude Code CLI evidence (actual model `glm-5.3`). This round's real Codex CLI attempt was blocked by read policy, not reported as passing. Local changes are not published.

- An empty-directory todo trial showed the agent skipping project selection and writing a static page. `soloweave` and `feature-workflow` are now `0.2.2` with an explicit new-project routing gate; `project-setup` is `0.3.1` and names empty directories in its description. Live retesting is pending.
- `project-setup` is now `0.3.0`: it discovers requirements and project type before comparing independently combinable choices by relevant dimension. Its guide offers decision questions and examples without limiting recommendations. SkillHub packages now include required reference files.
- Following a two-window Codex trial, require a clear status and handoff when a new project is blocked, and resume known unfinished work before asking for a new direction. The three affected Skills are now `0.2.1`.
- Added six new root Skills, including an ordinary-development entry and a debugging workflow; the Go CLI keeps its published legacy snapshot.
- Changed the default guide to installing Skills once and then asking for development normally. Added bilingual usage and acceptance docs, a migration design, and an ADR.
- Local SkillHub packaging now reads six root Skills and CI checks them. Codex files installed by an AI agent from public GitHub matched the source; observable automatic invocation, the revised two-window handoff, and SkillHub updates remain to be tested.

## 0.1.1-preview (released)

- Added CLI-free planning, development, handoff, and quality-review paths to the five Skills, with a clear boundary between manual records and CLI checks.
- Added bilingual Skills-only guidance and clearer explanations of the two usage paths and CLI commands.
- Generate version-specific bilingual instructions inside the Windows ZIP and record a checkpoint before the quick-start project check.
- Fix the GitHub download link in the entry Skill.

## 0.1.0-preview (released)

- Initial local preview with project configuration, architecture approval and ADRs, continuity commands, Skill installation adapters for three platforms, five core Skills, and project checks.
- Added Chinese getting-started documentation, a repository directory guide, local build instructions, and English translations of public documentation.
- Fixed false stale-checkpoint warnings after committing unchanged contents and changed the Go module path to the GitHub repository address.
- Added a ready-to-run Windows x64 ZIP and a draft Release workflow, with download-first instructions for users.
