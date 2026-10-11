# SoloWeave Skills

[简体中文](../skills.md)

Root [`skills/`](../../skills/) is the only maintained source. Each directory contains SKILL.md and any required resources, including SkillHub publication extensions. Users install the Skills directly.

| Skill | Source version | Purpose |
| --- | --- | --- |
| [soloweave](../../skills/soloweave/SKILL.md) | 0.2.4 | Entry, project inspection, new/existing project routing. |
| [project-setup](../../skills/project-setup/SKILL.md) | 0.3.2 | Requirements, dimensional selection, developer-confirmed decisions. |
| [feature-workflow](../../skills/feature-workflow/SKILL.md) | 0.2.4 | Existing-project implementation, reuse, validation, records. |
| [quality-review](../../skills/quality-review/SKILL.md) | 0.2.1 | Actual changes, decisions, verification results. |
| [project-continuity](../../skills/project-continuity/SKILL.md) | 0.2.2 | Resumption and checkpoints after meaningful work or blockers. |

Each Skill can be installed independently. Updating does not automatically uninstall an older debugging Skill; inspect its directory and user edits before deciding to remove it. Its essential workflow and existing project conventions remain usable if other Skills are absent. Automatic invocation is client-dependent.

## Installation

Follow the [standard installer documentation](https://github.com/vercel-labs/skills):

```sh
npx skills@latest add ShiYuandi/SoloWeave
```

Use `--skill name` for a single Skill, or ask an AI assistant to inspect existing files and install complete directories while preserving user edits. Do not skip confirmation on a modified installation.

Ask for work normally, checking that new projects clarify key decisions before implementation. Explicit invocation is a fallback and must be reported separately from automatic invocation. See the [workflow](skill-only-workflow.md) and [independent-install checks](../independent-skills-acceptance.md). The [selection guide](../../skills/project-setup/references/selection-guide.md) provides nonbinding examples.

## SkillHub publication status

Local source now contains five Skills; entry and feature development are 0.2.4. The developer has authorized committing and pushing this revision to GitHub; SkillHub synchronization remains pending. GitHub installations and platform downloads still reflect their actual online versions. The earlier six-Skill submission remains a [historical snapshot](../release-status.md), not publication proof for this revision. See the [page links](../skills.md#skillhub-发布状态).

Versions above identify source, not completed SkillHub review. See [delivery status](../release-status.md). The packaging script produces five ZIPs and SHA256SUMS in ignored `dist/skills/`. See [releasing](releasing.md).
