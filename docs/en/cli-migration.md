# Migrating from the legacy Go CLI

[简体中文](../cli-migration.md)

Current SoloWeave source retires the Go CLI. Install the six Skills directly; no executable or replacement CLI is required.

The historical [v0.1.1-preview Release](https://github.com/ShiYuandi/SoloWeave/releases/tag/v0.1.1-preview) retains its original download and tagged source. Its five embedded Skills are frozen snapshots. This document does not imply that an online Release was changed.

1. Back up the target project and verify its README, YAML contract, approval digest, ADRs, handoff, and actual implementation.
2. Preserve valid goals, entry points, and decisions in existing Markdown or PROJECT.md. Keep confirmation history; new architecture requires new approval.
3. Update STATUS.md and HANDOFF.md with current work, real checks, blockers, and next steps.
4. Inspect project and global Skills before installing. Preserve user edits and merge deliberately; the old CLI installer does not upgrade the new Skills.
5. Archive old machine files only after confirming no workflow still needs them. Migration does not automatically delete files or installations in other projects.

Former setup and approval commands become a developer-confirmed discussion recorded in PROJECT and ADRs. Installation uses the standard Skills installer or an AI assistant. Checkpoints and resumption use Markdown. Existing project checks replace diagnostic commands; there is no equivalent machine approval or stale-check guarantee. Each SKILL.md carries its own version.

This repository's obsolete machine records are preserved in the [historical directory](../archive/go-cli/README.md). See the archived [configuration](../archive/go-cli/en/configuration.md) and [continuity](../archive/go-cli/en/project-continuity.md) guides only for historical reference.
