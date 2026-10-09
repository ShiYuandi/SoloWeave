# Agent Skills and platform files

[简体中文](../skills.md)

The five canonical `SKILL.md` files are in `internal/bundle/assets/skills/` and are embedded in the Go binary. They follow the [Agent Skills specification](https://agentskills.io/specification): a short `name` and `description` help an agent select the right skill, while the body describes its workflow.

| Skill | Purpose |
| --- | --- |
| `soloweave` | Entry point and task routing |
| `project-setup` | Planning, stack selection, and architecture approval |
| `feature-workflow` | Inspecting and reusing existing code, then implementing a feature |
| `project-continuity` | Checkpoints and context recovery |
| `quality-review` | Running real checks and reporting results accurately |

`install --dry-run` previews writes without modifying files. Codex and Cursor load project skills from `.agents/skills/`; Claude Code uses `.claude/skills/`. The installer also writes concise platform rules and records file hashes in `.soloweave/installation.json`. On later runs, changes to managed files or existing user files are reported as conflicts. Resolve these explicitly; SoloWeave does not silently overwrite them.

Installation paths were checked against the [Codex](https://learn.chatgpt.com/docs/build-skills), [Claude Code](https://code.claude.com/docs/en/skills), and [Cursor](https://cursor.com/docs/skills) documentation. Live Claude Code and Cursor clients have not been tested yet.
