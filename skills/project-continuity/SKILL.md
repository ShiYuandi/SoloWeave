---
name: project-continuity
description: 在软件项目开始新会话、完成重要任务、遇到阻塞或切换 AI 账号与 Agent 时使用；核对真实状态并维护可供下一位 Agent 接手的项目记录。
slug: shiyuandi-soloweave-project-continuity
version: 0.2.0
displayName: SoloWeave 项目交接
summary: 自动保存真实进度、验证和下一步，让新会话从项目文件接手。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 项目连续性

## 接手时

读取项目已有的目标、决策、状态和交接文档；优先查看 `.soloweave/context/PROJECT.md`、`STATUS.md`、`HANDOFF.md`、`CHANGES.md` 与相关 ADR。核对当前分支、HEAD、未提交改动及交接中提到的源文件。没有 Git 或无法读取时写“Git 不可用”。记录与实际文件不一致时指出过期内容，并以代码与当前文件为依据继续工作；不能依赖旧聊天记录。

## 检查点

完成有意义的任务、出现阻塞或准备换账号时，先检查本轮实际修改和验证，再更新项目已有的等价文档；没有时使用 `.soloweave/context/` 下的四个 Markdown 文件：

- `PROJECT.md` 保存稳定目标、入口和已确认 ADR 链接，避免每次重写。
- `STATUS.md` 保存已完成、当前任务、阻塞与下一步。
- `CHANGES.md` 记录重要改动及文件；不要复制完整 Git 日志。
- `HANDOFF.md` 简短说明当前任务、最近修改与文件、实际验证命令及结果、未解决问题、下一步、核对时间和可取得的 Git 状态。未运行的检查写“未运行”。

只记录经过核对的事实，不写入凭据、令牌或私钥。纯 Skills 流程不生成 CLI 的 `checkpoint.json`、`approval.json` 或安装记录，也不宣称完成机器过期检查。小型无持续影响的改动无需扩大成冗长日志，但应让下一次接手能识别重要进展。
