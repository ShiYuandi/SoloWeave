# SoloWeave Skills

[简体中文](../skills.md)

The new Skill source lives in the repository's [`skills/`](../../skills/) directory. Each `SKILL.md` has the name, description, and workflow required by the [Agent Skills specification](https://agentskills.io/specification), plus publication fields used by SkillHub. This is the source edited and packaged from now on. `internal/bundle/assets/skills/` is the legacy snapshot embedded in Go CLI v0.1.1-preview and does not receive new Skill updates automatically.

| Skill | When it applies |
| --- | --- |
| [`soloweave`](../../skills/soloweave/SKILL.md) | Entry for ordinary development requests: inspect the project, verify records, and route work. |
| [`project-setup`](../../skills/project-setup/SKILL.md) | New or existing project setup; compare independently combinable technologies by relevant dimension and confirm important choices. |
| [`feature-workflow`](../../skills/feature-workflow/SKILL.md) | Feature work and ordinary code changes, after looking for reusable code. |
| [`systematic-debugging`](../../skills/systematic-debugging/SKILL.md) | Reproduce, locate, and fix bugs or failing tests. |
| [`quality-review`](../../skills/quality-review/SKILL.md) | Review code, delivery quality, and actual verification results. |
| [`project-continuity`](../../skills/project-continuity/SKILL.md) | Resume work, complete important tasks, handle blockers, or switch agents. |

## Installation and invocation

Ask a compatible agent to install from the GitHub repository, or follow the [Skills CLI documentation](https://github.com/vercel-labs/skills) and run `npx skills@latest add ShiYuandi/SoloWeave` to select Skills. Its `--list` option lists discovered Skills. This run listed and installed all six from the **local repository** in a disposable project. Separately, the user asked an AI agent to install all six from public GitHub; the installed files matched the source at commit `5bc8f1e`. A direct standard-installer command against the public URL has not yet been run.

After installation, ask for development work normally. Each agent decides whether to invoke a Skill from its name and description, so automatic use requires live client testing and cannot be guaranteed. If necessary, explicitly ask it to “use SoloWeave to continue this project.” See the [Skills-only guide](skill-only-workflow.md) for the project file workflow.

The `project-setup` [selection guide](../../skills/project-setup/references/selection-guide.md) is read when needed. It supplies comparison questions and nonbinding examples; the agent may recommend any technology suitable for the requirements. Project types identify relevant decisions, while the developer confirms the final combination.

## SkillHub publication status

The five legacy Skills currently published on [SkillHub](https://skillhub.cn/) are [`soloweave`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave), [`project-setup`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-setup), [`feature-workflow`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-feature-workflow), [`project-continuity`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-continuity), and [`quality-review`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-quality-review). The published entry is version `0.1.3`; the other four are `0.1.2`. `systematic-debugging` is not published there. The pending repository version of `project-setup` is `0.3.0`; `soloweave`, `feature-workflow`, and `project-continuity` are `0.2.1`; `quality-review` is `0.2.0`, and the debugging Skill is `0.1.0`.

Running `scripts/package-skillhub.ps1` locally packages the six root Skills, including needed reference files, and SHA256SUMS under ignored `dist/skillhub/skills-first/`. Before publication, verify versions, slugs, ZIP contents, and SkillHub's import preview. Commit, push, and publish only when the owner explicitly requests them. Packaging does not update the legacy Go executable.
