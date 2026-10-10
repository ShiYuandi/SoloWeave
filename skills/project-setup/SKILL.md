---
name: project-setup
description: 创建或接入软件项目、确定技术栈和架构时使用；分析现有项目，向开发者展示必要选项，并在明确确认后记录决策和项目入口。
slug: shiyuandi-soloweave-project-setup
version: 0.2.0
displayName: SoloWeave 项目设置
summary: 在 AI 中分析项目、确认关键选型并留下可复用的项目记录。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目设置

## 识别现状

新项目先弄清目标、用户、交付范围和运行环境；已有项目先读 README、依赖与锁文件、目录、测试、现有规则和 ADR。不要为了安装 SoloWeave 顺手重建项目或改动既有架构。已有技术版本以实际依赖文件为准。

## 确认决定

只对影响实施的关键分歧提出少量可行选项，说明成本和取舍，推荐一项并等待开发者明确选择。已确定且获批准的架构不得在日常任务中自行改写。等待确认期间，可以继续不依赖该决定的调查或准备工作。

确认后优先更新项目已有的等价文档；需要新文件时，创建 `.soloweave/context/PROJECT.md`，写明项目目标、运行入口、主要模块、技术选型及来源。在 `.soloweave/decisions/ADR-NNNN.md` 记录日期、问题、选项、决定、理由和确认事实，并从 `PROJECT.md` 链接。未确认的方案标为“待确认”，不写成已批准 ADR。重大选型变化另建 ADR，并指出被替代的决定。

如果项目已有 CLI 管理的 `.soloweave/project.yaml` 与 `approval.json`，先核对其约束及是否被更新的 ADR 取代；纯 Skills 流程不手工伪造或修改这些机器批准文件。没有 CLI 时也能完成上述 Markdown 记录。需要持续交接时，按 `project-continuity` 建立最少的状态文档；若未安装该 Skill，至少记录当前任务与下一步。不要覆盖用户现有内容。
