---
name: feature-workflow
description: 实现功能、修改现有代码或完成普通开发任务时使用；核对项目约束，先找可复用实现，再开发、验证并更新重要进度。
slug: shiyuandi-soloweave-feature-workflow
version: 0.2.1
displayName: SoloWeave 功能开发
summary: 先读项目决策并检查已有实现，完成开发、真实验证和交接。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 功能开发

1. 明确本次需求与完成条件。读取相关项目规则、已确认的架构决策、当前状态和涉及的代码；交接信息须与实际文件核对。
2. 搜索相近实现、公共函数、组件、依赖和测试。复用职责相同且不会明显增加耦合的实现；否则说明为何新增代码更合适。
3. 对较大改动简要规划后实施。技术栈、模块边界或数据模型的重大变化先说明取舍并取得开发者确认，再按 `project-setup` 记录新的 ADR。普通实现不需要重复请求确认。
4. 运行与改动相关的项目检查。区分通过、失败、未运行和工具不可用；出现失败时调查并修复，不能把未运行项目测试当作通过。
5. 完成有意义的功能或遇到阻塞时，在结束回复前使用 `project-continuity` 更新状态、重要变更、验证和下一步；若未安装该 Skill，直接更新项目已有的交接文档。新建的持续开发项目至少留下能定位目标的 `PROJECT.md`，以及写明当前任务、验证结果、阻塞和下一步的 `STATUS.md`、`HANDOFF.md`；有实际重要改动时记录 `CHANGES.md`。构建或启动未完成也要如实留下检查点，不能只在聊天回复中说明。小型格式或文字修改可保持记录简短。

没有 Go CLI 时照常完成开发和 Markdown 交接，不运行也不声称通过 `soloweave check`。若项目没有 SoloWeave 文件，沿用现有规范；需要长期协作记录时再建立最少文件。
