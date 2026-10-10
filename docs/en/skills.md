# Agent Skills and platform files

[简体中文](../skills.md)

The five canonical `SKILL.md` files are in `internal/bundle/assets/skills/` and are embedded in the Go binary. They follow the [Agent Skills specification](https://agentskills.io/specification): a short `name` and `description` help an agent select the right skill, while the body describes its workflow.

| Skill | Purpose |
| --- | --- |
| [`soloweave`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave) | Entry point and task routing |
| [`project-setup`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-setup) | Planning, stack selection, and architecture approval |
| [`feature-workflow`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-feature-workflow) | Inspecting and reusing existing code, then implementing a feature |
| [`project-continuity`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-continuity) | Checkpoints and context recovery |
| [`quality-review`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-quality-review) | Running real checks and reporting results accurately |

## Ask your AI to install from SkillHub

Each public page offers a prompt you can send to your coding agent. Copy the [five-Skill installation prompt in the README](README.md#let-your-ai-install-the-skills) and ask the agent to follow the [SkillHub installation guide](https://skillhub.cn/install/skillhub.md) for each Skill. Afterward, check all five `SKILL.md` paths and versions and resolve any conflicts with existing or edited files.

SkillHub supplies the Skill files only. Without the CLI, the AI can manually record developer-approved ADRs, project status, and handoffs using the [Skills-only guide](skill-only-workflow.md). This does not produce CLI approval records, platform rules, or machine-checkable checkpoints. `check` and `doctor` do not count a SkillHub-only installation as a CLI-managed install. For CLI automation, download the ready-to-run Windows x64 program from [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases), then follow the README's setup and installation preview.

`install --dry-run` previews writes without modifying files. Codex and Cursor load project skills from `.agents/skills/`; Claude Code uses `.claude/skills/`. The installer also writes concise platform rules and records file hashes in `.soloweave/installation.json`. On later runs, changes to managed files or existing user files are reported as conflicts. Resolve these explicitly; SoloWeave does not silently overwrite them.

Installation paths were checked against the [Codex](https://learn.chatgpt.com/docs/build-skills), [Claude Code](https://code.claude.com/docs/en/skills), and [Cursor](https://cursor.com/docs/skills) documentation. Live Claude Code and Cursor clients have not been tested yet.

## Publishing on SkillHub

Here SkillHub means [skillhub.cn](https://skillhub.cn/). The five canonical skills include SkillHub's `slug`, `version`, and `displayName` fields in addition to the Agent Skills `name` and `description`. Changes to these sources reach target projects with the next CLI release.

The preferred route is SkillHub's **Import from GitHub** flow. Select `ShiYuandi/SoloWeave` and all five skills. Commit and push the intended source changes first, refresh the import list, then review each slug, display name, version, description, and archive before submitting. SkillHub reads the remote repository, not unpushed local files.

Local upload is a fallback. On Windows, run `scripts/package-skillhub.ps1` in PowerShell to validate the five directories and create individual ZIPs plus `SHA256SUMS` under the ignored `dist/skillhub/` directory. You can also publish each directory directly from `internal/bundle/assets/skills/<name>/`. Run SkillHub's `--dry-run` before submitting each skill. Its CLI guide asks Windows users to use WSL; the website offers an upload flow. Publishing requires a registered, verified account and is subject to review. See the [official guide](https://skillhub.cn/tutorials) and [publishing requirements](https://skillhub.cn/ai/release.md).

For an existing listing, the dashboard's Update form accepts a new ZIP. Keep the same `slug`, increment the version, and enter a changelog. The new version goes through security review; the public page may continue to show the previous version during review.

Publish all five skills together. The `soloweave` entry skill calls the other four. The development conventions do not require the CLI. To use `soloweave` commands, download the CLI separately from [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) and add it to `PATH`. Only a Windows x64 executable is currently delivered. A SkillHub installation alone does not make the CLI available.

The SkillHub website requires a numeric `X.Y.Z` version. The first published version of all five Skills was `0.1.0`, followed by the Chinese `0.1.1` update. On 2026-10-10, all five `0.1.2` public pages were checked and offered downloads; their bodies include the Skills-only path. The entry Skill's GitHub download link was fixed in `0.1.3`; its public page now shows the correct link and a `0.1.3` download, while the other four remain at `0.1.2`. The current `0.1.1-preview` Windows x64 CLI bundles the entry Skill at `0.1.3` and the other four at `0.1.2`. Skills and CLI have separate release channels. For updates, keep each `slug` stable, document CLI compatibility, rebuild and validate affected packages, then submit new versions. Check the public SkillHub pages for current availability.
