---
name: soloweave
description: 在 SoloWeave 项目中开始或继续工作时使用；根据任务需要，引导至项目规划、功能开发、项目交接或质量检查流程。
slug: shiyuandi-soloweave
version: 0.1.0
displayName: SoloWeave 项目入口
summary: 在 Codex、Claude Code 和 Cursor 中延续项目决策、开发任务与交接信息。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# SoloWeave

本 Skill 是五个 Skill 组成的套件入口。请同时安装 `project-setup`、`feature-workflow`、`project-continuity` 和 `quality-review`。这些流程还会使用独立的 `soloweave` 命令行程序；Windows x64 程序和安装说明见 https://github.com/ShiYuandi/SoloWeave/releases。如果命令行程序不可用，应说明哪些步骤无法执行，不要声称已完成。

读取 `.soloweave/project.yaml`，了解已批准的工程决策；读取 `.soloweave/context/HANDOFF.md`，了解当前任务。采信交接信息前，先核对相关源文件和 Git 状态。如果项目尚无工程配置，使用 `project-setup`；开发功能或修复问题时，使用 `feature-workflow`；暂停工作或切换 Agent 时，使用 `project-continuity`；需要验证时，使用 `quality-review`。

重大架构变更由开发者决定。不要把草稿配置当作已批准的配置。按任务规模安排流程；小改动无需执行全部步骤。
