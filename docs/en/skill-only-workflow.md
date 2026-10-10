# Use SoloWeave with Skills only

[简体中文](../skill-only-workflow.md)

This path needs the five SoloWeave Skills but no `soloweave.exe`. Your AI can follow the engineering conventions and keep approved decisions and handoffs in ordinary project files. It does not provide the CLI's automatic configuration checks, installation record, or stale-checkpoint detection.

## Get started

Use the [README prompt](README.md#let-your-ai-install-the-skills) to install all five Skills, then ask the AI to report the path and version of each `SKILL.md`. In your target project, say: “Follow SoloWeave's conventions. Inspect the existing code and project records first. Without the CLI, use the Skills-only workflow and tell me which CLI checks were not run.”

## Plan and approve decisions

The AI first reads the existing README, dependency files, code, and ADRs. It preserves established decisions. When a new record is needed, it writes the goal, users, main modules, technical choices, run commands, and open decisions in `.soloweave/context/PROJECT.md`. Technical choices remain “pending developer approval” until you explicitly confirm them.

Only after you confirm a decision should the AI create `.soloweave/decisions/ADR-NNNN.md` with the date, options and trade-offs, decision, rationale, and approval fact, then link it from `PROJECT.md`. A major change needs another confirmation and ADR. A manual ADR is useful to the next AI, but **it is not CLI approval**: do not hand-create `approval.json` or change `project.yaml` to `approved` to imitate the CLI.

## Develop, review, and hand off

Before editing, ask the AI to inspect existing implementations, reusable components, and tests. After editing, run relevant project tests and report each command and actual result. Mark checks that did not run as “Not run”; without the CLI, mark `soloweave check` as “Tool unavailable,” never as passed.

At a handoff, update the files under `.soloweave/context/`. Use equivalent existing records first instead of making duplicates:

| File | Minimum content |
| --- | --- |
| `PROJECT.md` | Goal, approved decisions and ADRs, key entry points, and run commands |
| `STATUS.md` | Completed, in progress, blocked, and next work |
| `CHANGES.md` | Date, important changes, and affected files |
| `HANDOFF.md` | Current task, recent changes, actual and unrun checks, issues, next step, review time, and Git state |

If Git is available, check and record the branch, HEAD, and working-tree changes; otherwise write “Git unavailable.” Keep passwords, tokens, and private keys out of handoff files. A new AI reads these records and then checks the source files and Git. If they disagree, prefer the actual files and flag the record as possibly stale. Skills-only mode does not generate `checkpoint.json` or automatically detect stale checkpoints.

## If you add the CLI later

On Windows x64, download the ready-to-run ZIP from [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases); no source build is needed. Review existing files, then follow the [README CLI path](README.md#quick-start) to initialize, explicitly approve, and preview installation. The CLI does not automatically convert a manual ADR into its `approval.json`. Resolve any conflicting files before writing over existing work.
