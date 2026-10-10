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

## Publishing on SkillHub

Here SkillHub means [skillhub.cn](https://skillhub.cn/). The five canonical skills include SkillHub's `slug`, `version`, and `displayName` fields in addition to the Agent Skills `name` and `description`. Changes to these sources reach target projects with the next CLI release.

The preferred route is SkillHub's **Import from GitHub** flow. Select `ShiYuandi/SoloWeave` and all five skills. Commit and push the intended source changes first, refresh the import list, then review each slug, display name, version, description, and archive before submitting. SkillHub reads the remote repository, not unpushed local files.

Local upload is a fallback. On Windows, run `scripts/package-skillhub.ps1` in PowerShell to validate the five directories and create individual ZIPs plus `SHA256SUMS` under the ignored `dist/skillhub/` directory. You can also publish each directory directly from `internal/bundle/assets/skills/<name>/`. Run SkillHub's `--dry-run` before submitting each skill. Its CLI guide asks Windows users to use WSL; the website offers an upload flow. Publishing requires a registered, verified account and is subject to review. See the [official guide](https://skillhub.cn/tutorials) and [publishing requirements](https://skillhub.cn/ai/release.md).

Publish all five skills together. The `soloweave` entry skill calls the other four. SkillHub installs skill files, but users must separately download the `soloweave` CLI from [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) and add it to `PATH`. Only a Windows x64 executable is currently delivered. A SkillHub installation alone does not make the CLI available.

For updates, keep each `slug` stable, update the five skill versions to match the compatible SoloWeave release, rebuild the packages, validate, and submit the new versions. Check the actual SkillHub review result before claiming a listing is live.
