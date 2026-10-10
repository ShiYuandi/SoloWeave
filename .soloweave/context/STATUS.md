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

## 未完成或待外部验证

- 标准安装器首次尝试遇到 npm 连接拒绝，重试后的本地发现及安装已通过；公开 GitHub 安装尚未验证。
- 辅助 `quick_validate.py` 缺少 PyYAML，未完成该工具的校验；仓库打包脚本已执行名称、字段、版本和 slug 检查。
- 新版 Skill 在 Codex、Claude Code、Cursor 的真实自动触发与跨客户端交接尚未实测；新版尚未推送 GitHub 或更新 SkillHub。
- 旧版 `checkpoint.json` 对当前工作区可能过期，不能据此称机器检查点有效。

## 下一步

按 `docs/skills-first-plan.md` 的下一轮顺序，先验证公开 GitHub 安装，再于真实客户端验收普通请求自动选择、重大决策确认、重要任务后的交接和新会话恢复。SkillHub 更新仍需另行明确要求。
