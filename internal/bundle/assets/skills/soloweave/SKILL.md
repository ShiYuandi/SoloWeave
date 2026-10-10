---
name: soloweave
description: 在 SoloWeave 项目中开始或继续工作时使用；根据任务需要，引导至项目规划、功能开发、项目交接或质量检查流程。
slug: shiyuandi-soloweave
version: 0.1.2
displayName: SoloWeave 项目入口
summary: 在 Codex、Claude Code 和 Cursor 中按工程规范持续开发；不安装 CLI 也可记录决策与交接。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# SoloWeave

本 Skill 是五个 Skill 组成的套件入口。请同时安装 `project-setup`、`feature-workflow`、`project-continuity` 和 `quality-review`。先检查当前项目是否已有 SoloWeave 文件、已批准的决策和可用的 `soloweave` 命令行程序。CLI 是可选的本地自动化工具；没有 CLI 时使用各 Skill 的人工文件流程，不要求开发者为了遵守规范而先安装程序。CLI 的下载见 https://github.com/ShiYuandi/SoloWeave/releases。

优先读取已有的 `.soloweave/project.yaml`、`.soloweave/context/PROJECT.md`、决策记录和 `HANDOFF.md`；不存在时查看项目 README、依赖文件和代码，不把缺失文件当作已批准的架构。采信交接信息前，核对相关源文件和 Git 状态。如果项目尚无工程配置，使用 `project-setup`；开发功能或修复问题时，使用 `feature-workflow`；暂停工作或切换 Agent 时，使用 `project-continuity`；需要验证时，使用 `quality-review`。

重大架构变更由开发者决定。不要把草稿或未经确认的口头建议当作已批准的配置。没有 CLI 时，不要生成或冒充 CLI 的 `approval.json`、`checkpoint.json`，也不要声称执行过 `soloweave check`。按任务规模安排流程；小改动无需执行全部步骤。
