---
name: soloweave
description: 在软件项目中开始或继续开发、规划、修复、审查或接手工作时使用；自动读取项目事实，选择 SoloWeave 的适用工作流，并在重要节点维护可交接的上下文。
slug: shiyuandi-soloweave
version: 0.2.1
displayName: SoloWeave 开发入口
summary: 正常提出开发需求，AI 主动核对项目、选择工作流并维护交接。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# SoloWeave 开发入口

用户正常描述开发需求时，主动判断是否适用本技能。不要要求用户先输入 SoloWeave 命令。本技能只负责进入项目和选择工作流；具体任务按相应 Skill 完成。

## 进入项目

1. 查找项目规则、README、相关源码和测试。若有 `.soloweave/context/`，读取 `PROJECT.md`、`STATUS.md`、`HANDOFF.md` 与相关 ADR；若有 `.soloweave/project.yaml`，核对其状态、批准记录及后续 ADR，未经确认的选型仍是草稿，已被替代的决定不再约束新工作。
2. 核对交接中的文件、分支、HEAD 和未提交改动。没有 Git 或无法读取时如实说明；记录与实际文件矛盾时以实际文件为准。
3. 没有 SoloWeave 记录的现有项目，先沿用已有文档。开始持续开发或需要跨会话交接时，再按 `project-setup` 建立最少的项目记录。不要为一次小型文字修改创建整套文档。

## 选择工作流

- 新项目、接入项目或重要技术选型：`project-setup`。
- 功能开发和一般代码改动：`feature-workflow`。
- 行为错误、测试失败或性能回退：`systematic-debugging`。
- 代码审查、交付检查：`quality-review`。
- 恢复旧任务、完成重要节点或准备切换 Agent：`project-continuity`。

“继续这个项目”也是接手请求：先找交接中的下一步和已知阻塞，再核对代码与可运行的检查。缺少交接文件时，从项目事实重建当前状态；有明确的未完成工作就先推进它，不先让开发者从泛泛的方向列表中选择。只有目标仍无法判断、存在同等合理且影响较大的路线，或关键决定需要确认时再提问。未实际验证启动或构建时，不把项目称为“可运行”。

如果其他 Skill 尚未安装，直接按本入口的共同原则继续，并告知缺少的工作流；不要让用户为了继续任务而安装 Go CLI。重大架构决定先由开发者确认。完成有意义的工作后主动更新项目状态和交接；记录真实验证结果，未运行的检查不得写成通过。尊重项目既有规则与用户对提交、推送和发布的授权边界。
