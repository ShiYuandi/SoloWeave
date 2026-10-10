---
name: project-setup
description: 创建新的 SoloWeave 项目或接入现有项目时使用；帮助确认技术选型并记录开发者批准的架构决策。
slug: shiyuandi-soloweave-project-setup
version: 0.1.2
displayName: SoloWeave 项目规划
summary: 与开发者确认技术选型并记录决策；CLI 可选，支持仅用项目文档规划。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目规划

接入现有项目时，先查看源代码、依赖清单和测试，再建议工程配置。不要借接入之机迁移项目架构。创建新项目时，先明确产品目标和最小可行产品，再提出少量合适的技术栈与架构选项，说明各自取舍。最终选型由开发者决定。

如果 CLI 可用，使用 `soloweave init` 创建草稿版 `.soloweave/project.yaml`。与开发者核对配置后，再由开发者确认并运行 `soloweave approve`，生成 ADR。不要自行批准草稿，也不要擅自更改此前已批准的选项。

如果没有 CLI，先查找项目已有的架构文档和 ADR，避免覆盖。没有合适记录时，在 `.soloweave/context/PROJECT.md` 写明目标、项目类型、技术选项及取舍、主要模块、运行方式，并把尚未确认的选择标为“待开发者确认”。向开发者展示拟采用的方案；**只有开发者明确确认后**，才在 `.soloweave/decisions/ADR-NNNN.md` 记录日期、选项与取舍、最终决定、理由和确认事实，并在 `PROJECT.md` 链接该 ADR。架构改变时重复确认并新增 ADR。仅人工记录不等于 CLI 批准：不要手工创建 `.soloweave/approval.json`，也不要仅凭文档把 `project.yaml` 的状态改为 `approved`。未来接入 CLI 时另行运行其初始化与批准流程。

具体技术版本以项目依赖文件和锁文件为准。无法确认技术选型时，保留待确认状态并继续其他不依赖该决定的工作。
