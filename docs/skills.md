# SoloWeave Skills

[English](en/skills.md)

新版 Skill 源码位于仓库根目录 [`skills/`](../skills/)；每个目录都有 `SKILL.md`，包含 [Agent Skills 规范](https://agentskills.io/specification)要求的名称、描述和工作流，并保留 SkillHub 发布所需的扩展字段。这是今后编辑和打包的来源。Go CLI 内的 `internal/bundle/assets/skills/` 是 v0.1.1-preview 的旧版快照，不随新版 Skill 自动更新。

| Skill | 何时使用 |
| --- | --- |
| [`soloweave`](../skills/soloweave/SKILL.md) | 普通开发请求的入口：识别项目、核对记录、选择工作流。 |
| [`project-setup`](../skills/project-setup/SKILL.md) | 新项目、接入现有项目、确认关键技术与架构决策。 |
| [`feature-workflow`](../skills/feature-workflow/SKILL.md) | 实现功能和普通代码改动，先查找可复用实现。 |
| [`systematic-debugging`](../skills/systematic-debugging/SKILL.md) | 复现、定位并修复错误或失败测试。 |
| [`quality-review`](../skills/quality-review/SKILL.md) | 审查代码、交付质量和真实验证结果。 |
| [`project-continuity`](../skills/project-continuity/SKILL.md) | 新会话接手、重要任务结束、阻塞或切换 AI 时维护交接。 |

## 安装与调用

可以让支持 Agent Skills 的工具从 GitHub 仓库安装，或按 [Skills CLI 文档](https://github.com/vercel-labs/skills)使用 `npx skills@latest add ShiYuandi/SoloWeave` 选择 Skill。CLI 的 `--list` 可列出它发现的 Skill。本轮已从**本地仓库**列出六个并安装到隔离项目；尚未从公开 GitHub 地址运行该命令，在线安装结果以实际运行记录为准。

安装后直接提出开发需求。Agent 会依据 Skill 的名称和描述决定是否调用，因此自动触发需要真实客户端验证，不能保证每次发生。若未触发，可明确说“使用 SoloWeave 继续这个项目”。无 CLI 的完整项目文件流程见[使用指南](skill-only-workflow.md)。

## SkillHub 发布状态

[SkillHub](https://skillhub.cn/) 目前公开的五个旧版 Skill 分别是 [`soloweave`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave)、[`project-setup`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-setup)、[`feature-workflow`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-feature-workflow)、[`project-continuity`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-continuity)、[`quality-review`](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-quality-review)。入口公开版本为 `0.1.3`，其余四项为 `0.1.2`；新增的 `systematic-debugging` 尚未上架。根目录五个旧名称的新稿标为 `0.2.0`，调试 Skill 标为 `0.1.0`，均不能在发布前称为 SkillHub 已更新。

本地运行 `scripts/package-skillhub.ps1` 会从根目录六个 Skill 生成 ZIP 和 SHA256SUMS，输出到忽略的 `dist/skillhub/skills-first/`。发布前核对版本、slug、压缩包内容及 SkillHub 导入预览；用户明确要求后再提交、推送和发布。脚本不会更新旧版 Go 程序中的 Skill。
