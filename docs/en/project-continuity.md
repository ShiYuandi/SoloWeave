# Project continuity

[简体中文](../project-continuity.md)

SoloWeave keeps recoverable project context under `.soloweave/context/`:

| File | Purpose | Update when |
| --- | --- | --- |
| `PROJECT.md` | Stable goals, modules, run commands, and entry points | Important project facts change |
| `STATUS.md` | Completed, active, blocked, and planned work | Task status changes |
| `HANDOFF.md` | Current task, recent changes, actual checks, risks, and next step | At a checkpoint or handoff |
| `CHANGES.md` | Important completed changes | After a feature or significant fix |
| `tasks/` | Detailed task records | A task needs durable notes |

`context checkpoint` writes `HANDOFF.md` and the CLI-owned `checkpoint.json`. The agent or developer provides the business summary and actual verification results; when Git is available, the CLI records the branch, HEAD, working tree state, and a digest of project file contents. `context check` warns when the branch or project contents change. Committing identical contents does not invalidate a new checkpoint. Older checkpoints still use the previous Git-state comparison, so create a new checkpoint to adopt this behavior. Without Git, the CLI reports that Git verification is unavailable. The CLI cannot recover unrecorded work after the last checkpoint or infer business intent from code.

A new agent should run `context resume`, read the approved configuration and task notes, then inspect the actual code and Git state. Treat discrepancies as a sign that the handoff may be stale. Never put API keys, tokens, passwords, private keys, or other real secrets in these files.

When using only the Skills without the CLI, follow the [Skills-only guide](skill-only-workflow.md) to maintain the same project records manually. Check the real files and Git state when taking over. A manual handoff does not create `checkpoint.json` or pass the CLI's automatic stale-checkpoint check.
