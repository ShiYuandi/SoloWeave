---
name: project-continuity
description: 记录 SoloWeave 项目检查点、交接给其他 AI 账号或编程 Agent，以及在聊天上下文丢失后继续开发时使用。
slug: shiyuandi-soloweave-project-continuity
version: 0.1.0
displayName: SoloWeave 项目交接
summary: 记录真实开发进度，让新会话或不同编程 Agent 能接续项目。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目交接

继续开发时，读取 `.soloweave/project.yaml`，以及 `.soloweave/context/` 下的 `PROJECT.md`、`STATUS.md` 和 `HANDOFF.md`；如果命令行程序可用，运行 `soloweave context resume`。核对 Git 状态和相关源文件。如果它们与交接记录不一致，应指出交接记录已过期，并以实际文件为依据。

在有意义的检查点，更新 `STATUS.md` 中已完成、进行中、受阻和计划中的工作；将重要的已完成改动写入 `CHANGES.md`。运行 `soloweave context checkpoint`，提供简明的任务摘要、下一步，以及实际执行过的验证结果。让 `HANDOFF.md` 足够简短，使新 Agent 无需旧聊天记录也能继续工作。不要在上下文文件中写入凭据、令牌、私钥或真实密钥值。如果 Git 不可用，应明确说明。
