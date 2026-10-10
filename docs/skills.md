# Agent Skill 与平台文件

[English](en/skills.md)

五个标准 `SKILL.md` 源文件位于 `internal/bundle/assets/skills/`，构建时嵌入 Go 程序。它们遵循 [Agent Skills 规范](https://agentskills.io/specification)：简短的 `name`、`description` 帮助 Agent 选择技能，正文说明工作流程。

| Skill | 用途 |
| --- | --- |
| [`soloweave`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave) | 总入口与任务路由 |
| [`project-setup`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-setup) | 规划、技术选型、架构确认 |
| [`feature-workflow`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-feature-workflow) | 检查现有代码、复用和实现功能 |
| [`project-continuity`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-continuity) | 检查点与上下文恢复 |
| [`quality-review`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-quality-review) | 执行真实检查并如实报告结果 |

`install --dry-run` 只预览写入，不修改文件。Codex 与 Cursor 从 `.agents/skills/` 加载项目 Skill；Claude Code 使用 `.claude/skills/`。安装器还会写入相应平台的简短规则，并把文件摘要记录在 `.soloweave/installation.json`。再次安装时，已被修改的受管文件或用户已有文件会产生冲突，需要人工处理；安装器不会静默覆盖。

安装路径已按 [Codex](https://learn.chatgpt.com/docs/build-skills)、[Claude Code](https://code.claude.com/docs/en/skills) 和 [Cursor](https://cursor.com/docs/skills) 文档核对。真实 Claude Code/Cursor 客户端尚未完成实测。

## 在 SkillHub 发布

这里的 SkillHub 指 [skillhub.cn](https://skillhub.cn/)，不是同名的其他平台。五个源 Skill 的 frontmatter 除 Agent Skills 的 `name`、`description` 外，还包含 SkillHub 所需的 `slug`、`version`、`displayName`，以及简介和 MIT 许可证。修改源文件后，CLI 安装到目标项目的内容也会随下一个程序版本更新。

首选在 SkillHub 网页的「从 GitHub 导入」中选择 `ShiYuandi/SoloWeave`，一次勾选五个源 Skill。**先把准备发布的改动提交并推送到 GitHub**，再在导入页点「刷新」，核对每个 Skill 的 slug、显示名称、版本、描述和归档内容，然后提交审核。SkillHub 会读取远端仓库内容，不会读取本地尚未推送的文件。

本地上传是备用路径。在 Windows PowerShell 中运行 `scripts/package-skillhub.ps1`，会校验五个源目录并在忽略的 `dist/skillhub/` 生成各自的 ZIP 与 `SHA256SUMS`。各 Skill 也可以直接从 `internal/bundle/assets/skills/<名称>/` 目录发布。发布前可按 [SkillHub 发布规范](https://skillhub.cn/ai/release.md)逐个进行 `--dry-run`；Windows 用户如选择 SkillHub CLI，请使用 WSL，或在网页端上传。平台要求发布者注册并完成实名认证；正式提交后还要等待审核。[官方教程](https://skillhub.cn/tutorials)

更新已发布的 Skill 时，管理页的「更新」表单支持上传新 ZIP。保留原 `slug`，递增版本号并填写变更说明；提交后新版本重新进入安全审核，审核期间公开页仍可能展示旧版内容。

发布时应提供五个 Skill：入口 `soloweave`、`project-setup`、`feature-workflow`、`project-continuity`、`quality-review`。入口会按任务调用其余四个，因此只安装入口并不构成完整流程。此外，SkillHub 安装的是 Skill 文件，`soloweave` 命令行程序仍需用户从 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) 单独下载并配置 `PATH`。当前只提供 Windows x64 程序；其他系统的可运行程序尚未交付。发布到 SkillHub 不等于 CLI 自动可用。

SkillHub 网页实际要求数字 `X.Y.Z` 版本。首发的五个 Skill 为 `0.1.0`；中文概述和正文对应源文件版本 `0.1.1`，已提交 SkillHub 安全审核。两版均配合 SoloWeave CLI `0.1.0-preview` Windows x64 下载版使用；Skill 与 CLI 属于不同发布渠道。后续更新要同时调整五个 Skill 的版本，说明兼容的程序版本；保留各自的 `slug`，重新生成包、校验并提交新版本。上架状态以 SkillHub 实际审核结果为准。
