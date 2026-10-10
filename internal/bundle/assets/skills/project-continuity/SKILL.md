---
name: project-continuity
description: 记录 SoloWeave 项目检查点、交接给其他 AI 账号或编程 Agent，以及在聊天上下文丢失后继续开发时使用。
slug: shiyuandi-soloweave-project-continuity
version: 0.1.2
displayName: SoloWeave 项目交接
summary: 记录真实进度、验证与下一步；没有 CLI 也能用项目文档交接。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目交接

继续开发时，读取现有的 `.soloweave/project.yaml`（如有）、`.soloweave/context/` 下的 `PROJECT.md`、`STATUS.md`、`HANDOFF.md` 和相关 ADR；如果 CLI 可用，运行 `soloweave context resume`。核对 Git 状态和相关源文件。如果它们与交接记录不一致，应指出交接记录可能过期，并以实际文件为依据；没有 CLI 时不能声称已通过机器检查点校验。

在有意义的检查点，更新 `STATUS.md` 中已完成、进行中、受阻和计划中的工作；将重要的已完成改动写入 `CHANGES.md`。CLI 可用时运行 `soloweave context checkpoint`，提供简明的任务摘要、下一步和实际验证结果。

没有 CLI 时，先保留并更新项目已有的同类文档；缺失时创建 `.soloweave/context/PROJECT.md`、`STATUS.md`、`HANDOFF.md`、`CHANGES.md`。`PROJECT.md` 写项目目标、已确认决策及入口；`HANDOFF.md` 简短记录当前任务、最近改动与文件、实际运行的验证及结果（未运行则写“未运行”）、未解决问题、下一步和核对时间。Git 可用时记录分支、HEAD 与工作区变化；不可用时写“Git 不可用”。此路径不生成 `checkpoint.json`，也不具备 CLI 的自动过期检测；下一个 Agent 必须再次核对文件与 Git。不要在上下文文件中写入凭据、令牌、私钥或真实密钥值。
