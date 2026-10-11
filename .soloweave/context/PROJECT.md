# SoloWeave 项目

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 开发 Skills，帮助开发者确认工程决策、复用已有实现，并将进度和交接信息留在项目文件中。根目录 `skills/` 是六个可独立安装的 Skill 的唯一当前来源；类型与参考指南用于发现选型维度，不限制 AI 可以推荐的技术。

开发者于 2026-10-11 确认实施 Go 退出计划。当前源码已移除 Go CLI、旧内置 Skills 和程序构建发布链；维护者使用 PowerShell 7 校验和打包，普通用户直接安装 Skills。旧配置、批准与机器检查点归档在 `docs/archive/go-cli/`，不再约束当前工程。

当前方向见 `docs/skills-first-design.md`、`docs/go-retirement-plan.md` 和 `.soloweave/decisions/ADR-0003.md`。早期设计和计划保留为历史；使用说明见 README 和 `docs/skills.md`，目录见 `docs/directory-structure.md`。

项目记录以 Markdown 事实和 ADR 为准，不能伪称机器校验或真实客户端测试已通过。发布流程见 `docs/releasing.md`；只有开发者明确要求才提交、推送或公开发布。