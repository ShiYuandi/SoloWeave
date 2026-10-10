---
name: feature-workflow
description: 在 SoloWeave 项目中开发功能或修复问题时使用；遵守已批准的架构，并优先检查可复用的现有代码。
slug: shiyuandi-soloweave-feature-workflow
version: 0.1.2
displayName: SoloWeave 功能开发
summary: 在已确认的架构下检查现有代码、复用合适实现并完成需求或修复。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 功能开发

先读取已有的工程配置或 `PROJECT.md`、ADR 和当前交接记录；没有 CLI 时以开发者明确确认的 ADR 为决策依据，未确认的选型不能视为既定架构。设计改动前，检查相关模块、现有组件、工具方法、依赖和测试。有合适的实现就复用；如果复用会增加耦合或模糊职责，则编写新代码。将改动控制在满足需求所需的范围内。

较大任务先简要说明实施计划。完成修改后运行相关验证，并如实报告实际执行的检查；CLI 不可用不妨碍运行项目自己的测试。如果任务涉及变更已批准的技术或架构决策，先取得开发者确认，再按 `project-setup` 记录新的 ADR。在有意义的检查点使用 `project-continuity` 更新项目状态和交接信息。
