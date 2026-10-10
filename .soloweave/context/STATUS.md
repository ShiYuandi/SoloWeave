# 当前状态

核对日期：2026-10-10。

## 已完成

- 用户确认 SoloWeave 改为 Skills 优先、普通开发请求由 AI 主动使用、项目上下文在重要节点自动维护、Go CLI 逐步退出。方向记录于 `docs/skills-first-design.md`、`docs/skills-first-plan.md` 和 `ADR-0002.md`。
- 根目录 `skills/` 新增六个新版 Skill：入口、项目设置、功能开发、故障排查、质量检查和项目交接。五个原名称保留，Go CLI 的旧版嵌入资源未改。
- 中英文 README、Skills 使用说明、目录、贡献和验收文档改为 Skills 优先；旧版 CLI 的下载与功能仍有说明。
- `scripts/package-skillhub.ps1` 改从根目录打包六个 Skill；CI 增加打包校验步骤。
- 本地六个 ZIP 和 SHA256SUMS 核对通过，改动 Markdown 的相对文件链接与 `git diff --check` 通过。旧版 Go `test ./... -count=1`、`vet ./...`、`build ./...` 退出码均为 0。
- 标准安装器 `add . --list` 在本地识别六个 Skill；临时项目完成 Codex、Claude Code、Cursor 的安装，所有安装后的 `SKILL.md` 与源文件一致；Windows 默认模式也完成 Codex 六个 Skill 的安装。
- 在无 SoloWeave CLI、无 Git 的临时项目中，人工完成测试先失败、复用已有函数后两项检查通过的开发流程，并写入项目和交接记录；单独读取后重新运行测试仍通过。
- 推送前重新运行旧版 Go CLI 的 `test ./... -count=1`、`vet ./...`、`build ./...`，均以退出码 0 完成；远端 `master` 在推送前与本地 HEAD `5d8f22b` 一致。
- Skills 优先改动已以中文提交 `5bc8f1e` 推送至 GitHub `master`。用户提供的 Codex 双窗口试用显示：六个全局安装文件与仓库源文件一致；新项目写入了 `PROJECT.md` 和 ADR，但缺少状态、变更和交接记录。第二个窗口未接住已知启动阻塞，跨会话连续性验收未通过。
- 按用户确认的修正方案，仓库源码的 `soloweave`、`feature-workflow`、`project-continuity` 升为 `0.2.1`，要求受阻时完成状态与交接检查点，并从已知未完成工作恢复。六个 Skill 的本地打包、ZIP 内容与 SHA256SUMS 核对通过；真实客户端尚未用修订版重测。
- 用户指出项目设置的技术推荐过窄；已按经审核的设计与实施计划修订根目录 `project-setup` 至 `0.3.0`。新规则先识别需求和类型，再按适用维度讨论可组合选型；参考指南不限制 AI 技术选择。打包脚本改为包含所需参考文件。六个本地 ZIP、SHA256SUMS 与逐文件字节核对通过；标准安装器在隔离项目复制的 `project-setup` 两个文件与源码 SHA256 一致。真实新会话行为待复测。

## 未完成或待外部验证

- 标准安装器首次尝试遇到 npm 连接拒绝，重试后的本地发现及安装已通过；用户通过 AI 从公开 GitHub 完成全局安装并核对文件一致，但公开地址的标准安装器命令尚未直接验证。
- 辅助 `quick_validate.py` 缺少 PyYAML，未完成该工具的校验；仓库打包脚本已执行名称、字段、版本和 slug 检查。
- Codex 已有用户提供的真实双窗口行为样本，但对话摘要无法确认内部自动调用事件；该样本的跨会话交接失败。Claude Code、Cursor 真实客户端尚未实测；新版尚未更新 SkillHub。
- `project-setup` `0.3.0` 的模糊需求、逐项选型、表外技术、组合兼容性及批准边界尚未在真实 Agent 新会话复测。`quick_validate.py` 仍因缺少 PyYAML 无法运行；本地打包脚本和标准安装器文件核对已通过。
- 旧版 `checkpoint.json` 对当前工作区可能过期，不能据此称机器检查点有效。

## 下一步

用户已明确要求将本轮修改提交并推送到 GitHub，完成情况以 Git 记录和远端状态为准。之后将 `project-setup` `0.3.0` 与交接修订 `0.2.1` 安装到真实客户端新会话，按 `docs/agent-acceptance.md` 记录选型与跨会话接手结果；公开 GitHub 标准安装器、Claude Code/Cursor 与 SkillHub 更新另行验证。SkillHub 发布仍需另行明确要求。
