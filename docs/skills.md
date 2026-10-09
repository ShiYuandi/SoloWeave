# Agent Skill 与平台文件

[English](en/skills.md)

五个标准 `SKILL.md` 源文件位于 `internal/bundle/assets/skills/`，构建时嵌入 Go 程序。它们遵循 [Agent Skills 规范](https://agentskills.io/specification)：简短的 `name`、`description` 帮助 Agent 选择技能，正文说明工作流程。

| Skill | 用途 |
| --- | --- |
| `soloweave` | 总入口与任务路由 |
| `project-setup` | 规划、技术选型、架构确认 |
| `feature-workflow` | 检查现有代码、复用和实现功能 |
| `project-continuity` | 检查点与上下文恢复 |
| `quality-review` | 执行真实检查并如实报告结果 |

`install --dry-run` 只预览写入，不修改文件。Codex 与 Cursor 从 `.agents/skills/` 加载项目 Skill；Claude Code 使用 `.claude/skills/`。安装器还会写入相应平台的简短规则，并把文件摘要记录在 `.soloweave/installation.json`。再次安装时，已被修改的受管文件或用户已有文件会产生冲突，需要人工处理；安装器不会静默覆盖。

安装路径已按 [Codex](https://learn.chatgpt.com/docs/build-skills)、[Claude Code](https://code.claude.com/docs/en/skills) 和 [Cursor](https://cursor.com/docs/skills) 文档核对。真实 Claude Code/Cursor 客户端尚未完成实测。
