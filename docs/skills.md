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

## 让 AI 从 SkillHub 安装

五个公开详情页都提供“将提示词发送给你的 AI 安装”入口。可直接复制 [README 中的五个 Skill 安装提示词](../README.md#让-ai-安装-skills)，让当前 AI 编程工具按 [SkillHub 安装说明](https://skillhub.cn/install/skillhub.md)逐个安装。安装后应核对五个 `SKILL.md` 的位置与版本，遇到同名或已修改文件先处理冲突。

SkillHub 安装只提供 Skill 文件。没有 CLI 时，AI 可按[纯 Skills 使用指南](skill-only-workflow.md)人工记录经开发者确认的 ADR、项目状态和交接；这不生成 CLI 的批准记录、平台规则或机器检查点。`check`/`doctor` 读取 CLI 自己的安装记录，不会把单独的 SkillHub 安装视为 CLI 已安装。需要 CLI 的项目自动化时，先从 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) 获取 Windows x64 可运行程序，再按 README 的 CLI 步骤初始化和预览安装。

`install --dry-run` 只预览写入，不修改文件。Codex 与 Cursor 从 `.agents/skills/` 加载项目 Skill；Claude Code 使用 `.claude/skills/`。安装器还会写入相应平台的简短规则，并把文件摘要记录在 `.soloweave/installation.json`。再次安装时，已被修改的受管文件或用户已有文件会产生冲突，需要人工处理；安装器不会静默覆盖。

安装路径已按 [Codex](https://learn.chatgpt.com/docs/build-skills)、[Claude Code](https://code.claude.com/docs/en/skills) 和 [Cursor](https://cursor.com/docs/skills) 文档核对。真实 Claude Code/Cursor 客户端尚未完成实测。

## 在 SkillHub 发布

这里的 SkillHub 指 [skillhub.cn](https://skillhub.cn/)，不是同名的其他平台。五个源 Skill 的 frontmatter 除 Agent Skills 的 `name`、`description` 外，还包含 SkillHub 所需的 `slug`、`version`、`displayName`，以及简介和 MIT 许可证。修改源文件后，CLI 安装到目标项目的内容也会随下一个程序版本更新。

首选在 SkillHub 网页的「从 GitHub 导入」中选择 `ShiYuandi/SoloWeave`，一次勾选五个源 Skill。**先把准备发布的改动提交并推送到 GitHub**，再在导入页点「刷新」，核对每个 Skill 的 slug、显示名称、版本、描述和归档内容，然后提交审核。SkillHub 会读取远端仓库内容，不会读取本地尚未推送的文件。

本地上传是备用路径。在 Windows PowerShell 中运行 `scripts/package-skillhub.ps1`，会校验五个源目录并在忽略的 `dist/skillhub/` 生成各自的 ZIP 与 `SHA256SUMS`。各 Skill 也可以直接从 `internal/bundle/assets/skills/<名称>/` 目录发布。发布前可按 [SkillHub 发布规范](https://skillhub.cn/ai/release.md)逐个进行 `--dry-run`；Windows 用户如选择 SkillHub CLI，请使用 WSL，或在网页端上传。平台要求发布者注册并完成实名认证；正式提交后还要等待审核。[官方教程](https://skillhub.cn/tutorials)

更新已发布的 Skill 时，管理页的「更新」表单支持上传新 ZIP。保留原 `slug`，递增版本号并填写变更说明；提交后新版本重新进入安全审核，审核期间公开页仍可能展示旧版内容。

发布时应提供五个 Skill：入口 `soloweave`、`project-setup`、`feature-workflow`、`project-continuity`、`quality-review`。入口会按任务调用其余四个，因此只安装入口并不构成完整流程。只用开发规范时无需 CLI；若要使用 `soloweave` 命令，则需从 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) 单独下载并配置 `PATH`。当前只提供 Windows x64 程序；其他系统的可运行程序尚未交付。发布到 SkillHub 不等于 CLI 自动可用。

SkillHub 网页实际要求数字 `X.Y.Z` 版本。首发的五个 Skill 为 `0.1.0`；中文更新版为 `0.1.1`。2026-10-10 已核对五个 `0.1.2` 公开详情页，均可下载安装，正文包含纯 Skills 路线。入口 Skill 的 GitHub 下载链接有句号识别问题，源码已准备 `0.1.3` 修复版，公开状态须另行核对。当前 CLI `0.1.0-preview` Windows x64 下载版仍包含打包时的旧 Skill；Skill 与 CLI 属于不同发布渠道。后续更新要调整受影响 Skill 的版本，说明兼容的程序版本；保留各自的 `slug`，重新生成包、校验并提交。上架状态以 SkillHub 实际页面为准。
