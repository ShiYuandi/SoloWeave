# SoloWeave 项目

SoloWeave 是面向个人开发者与 1–5 人团队的 AI 开发 Skills，帮助开发者确认关键决策、复用已有实现，并将实际进度和交接信息保存在项目文件中。

根目录 `skills/` 是五项可独立安装 Skill 的唯一来源。用户安装后正常提出开发需求；维护者使用 PowerShell 7 检查、打包和维护 GitHub CI。类型和参考指南用于发现适用维度，不限制 AI 推荐的技术。

当前设计见 `docs/design.md`，实施和后续计划见 `docs/development-plan.md`，文档方向见 `.soloweave/decisions/ADR-0004.md`，五项职责边界见 `.soloweave/decisions/ADR-0005.md`。使用入口见 README、`docs/skills.md` 和 `docs/skill-only-workflow.md`。

项目记录、实际文件和验证证据共同用于接手，不把过期交接或未运行检查视为事实。GitHub 与 SkillHub 发布步骤见 `docs/releasing.md`；提交、推送和公开发布需开发者明确要求。
