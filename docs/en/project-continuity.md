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

`context checkpoint` writes `HANDOFF.md` and the CLI-owned `checkpoint.json`. The agent or developer provides the business summary and actual verification results; when Git is available, the CLI records the branch, HEAD, and working tree state. `context check` compares saved and current Git state and warns about obvious staleness. Without Git, it reports that Git verification is unavailable. The CLI cannot recover unrecorded work after the last checkpoint or infer business intent from code.

A new agent should run `context resume`, read the approved configuration and task notes, then inspect the actual code and Git state. Treat discrepancies as a sign that the handoff may be stale. Never put API keys, tokens, passwords, private keys, or other real secrets in these files.
