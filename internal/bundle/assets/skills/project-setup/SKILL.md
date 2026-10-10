---
name: project-setup
description: 创建新的 SoloWeave 项目或接入现有项目时使用；帮助确认技术选型并记录开发者批准的架构决策。
slug: shiyuandi-soloweave-project-setup
version: 0.1.1
displayName: SoloWeave 项目规划
summary: 为新项目或现有项目确认技术选型，并记录开发者批准的架构决策。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目规划

接入现有项目时，先查看源代码、依赖清单和测试，再建议工程配置。不要借接入之机迁移项目架构。创建新项目时，先明确产品目标和最小可行产品，再提出少量合适的技术栈与架构选项，说明各自取舍。最终选型由开发者决定。

使用 `soloweave init` 创建草稿版 `.soloweave/project.yaml`。与开发者核对配置后，再运行 `soloweave approve`，将决策记录为 ADR。不要自行批准草稿，也不要擅自更改此前已批准的选项。具体技术版本应记录在项目的依赖文件中，而不是 SoloWeave 工程配置中。
