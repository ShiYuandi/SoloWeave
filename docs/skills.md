# SoloWeave Skills

[English](en/skills.md)

根目录 [`skills/`](../skills/) 是五个 Skill 的唯一维护与发布来源。每个 SKILL.md 包含名称、描述、工作流和 SkillHub 扩展发布字段；所需参考资料随整个 Skill 目录安装。普通用户直接安装 Skills 即可使用。

| Skill | 当前源码版本 | 何时使用 |
| --- | --- | --- |
| [soloweave](../skills/soloweave/SKILL.md) | 0.2.4 | 普通任务入口：读取项目、区分新旧项目并路由。 |
| [project-setup](../skills/project-setup/SKILL.md) | 0.3.2 | 新项目、接入项目与重要选型；先澄清、分维度比较，再确认决定。 |
| [feature-workflow](../skills/feature-workflow/SKILL.md) | 0.2.4 | 已有项目的功能开发、复用、验证和记录。 |
| [quality-review](../skills/quality-review/SKILL.md) | 0.2.1 | 核对实际改动、项目决策和验证结果。 |
| [project-continuity](../skills/project-continuity/SKILL.md) | 0.2.2 | 新会话接手、重要节点或受阻时维护交接。 |

每个 Skill 可独立安装。升级不会自动卸载已安装的旧故障排查 Skill；先核对其目录和用户修改，再由用户决定移除。未安装其他 Skill 时，按当前 Skill 中的必要步骤或已有项目规范继续；入口可协调多个流程。是否自动触发取决于客户端，不能从“我会使用某 Skill”推断实际调用。

## 安装与使用

按 [标准安装器说明](https://github.com/vercel-labs/skills)运行：

```sh
npx skills@latest add ShiYuandi/SoloWeave
```

也可以请 AI 安装完整目录，先检查同名文件并保留用户修改，再报告路径和版本。安装单个 Skill 时可使用安装器的 `--skill 名称`。不应在有用户修改的目录中直接使用跳过确认参数。

安装后正常提出开发需求。新项目应先确认会影响形态的缺口，已有项目先核对现状。必要时明确点名 Skill；手动点名的结果须与自动触发分开记录。详见[工作流](skill-only-workflow.md)和[独立安装验收](independent-skills-acceptance.md)。

[选型判断指南](../skills/project-setup/references/selection-guide.md)只是非约束性参考，类型、技术示例和组合不限制 AI 推荐适合的路线。

## SkillHub 发布状态

当前五项已在 GitHub 提交 37c7062 中提供，并于 2026-10-12 完成公开安装核对。SkillHub 入口和功能开发 0.2.4 均已发布。项目设置 0.3.2、项目交接 0.2.2、质量检查 0.2.1 已发布；这五项已发布版本的公开包共六份源码资源均与仓库一致；平台新增 _meta.json 的 slug/version 另行核对。旧故障排查已下架，未永久删除；详细状态见[交付记录](release-status.md)。页面：[入口](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave)、[项目设置](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-setup)、[功能开发](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-feature-workflow)、[项目交接](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-project-continuity)、[质量检查](https://skillhub.cn/skills/user_38c0807e/shiyuandi-soloweave-quality-review)。

上表记录源码版本；SkillHub 状态为核对时快照，本轮实际进度见[交付状态](release-status.md)。维护者运行 `scripts/package-skillhub.ps1` 生成 `dist/skills/` 下的五个 ZIP 和 SHA256SUMS。步骤见[发布指南](releasing.md)。
