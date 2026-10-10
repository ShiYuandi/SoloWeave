# 项目交接

核对日期：2026-10-10。

## 当前任务

将 SoloWeave 从 CLI 为主的使用方式迁移为 Skills 优先、AI 主动使用的开发工作流。用户已确认方向并要求先写文档再实施。

## 当前事实

分支 `master`；本轮提交及远端同步结果以 Git 记录为准。设计、六个新版 Skill、脚本、CI、双语文档和交接改动构成本轮交付。最新公开的 GitHub Release 仍为 `v0.1.1-preview` 旧版 CLI；SkillHub 仍是五个旧版 Skill，尚无新版发布。本仓库 `project.yaml` 和 `approval.json` 仍代表旧版 Go 子工程；整体新方向见 `ADR-0002.md`。

## 最近修改

- `skills/`：六个新版 Skill，普通开发入口负责路由，其他 Skill 分别处理设置、功能、调试、质量与交接。
- `docs/skills-first-design.md`、`docs/skills-first-plan.md`：迁移设计和阶段计划。
- `README.md`、`docs/en/README.md`、使用、贡献与验收文档：Skills 安装和自然语言使用作为默认入口；CLI 列为旧版可选工具。
- `scripts/package-skillhub.ps1`、`.github/workflows/ci.yml`：从根目录打包六个 Skill 并加入 CI 校验。

## 实际验证

- `pwsh -File scripts/package-skillhub.ps1`：通过，生成六个 ZIP；逐包核对单个 `SKILL.md` 条目和 SHA256SUMS 通过。
- `.tools/go/bin/go.exe test ./... -count=1`、`vet ./...`、`build ./...`：退出码均为 0。Go 运行时另有无法写入本机遥测令牌的提示，不影响上述退出码。
- `git diff --check`：通过。改动 Markdown 的相对文件链接检查：通过。
- `quick_validate.py`：未完成，本机缺少 PyYAML。标准安装器本地 `--list`：未完成，连接 npm 注册表时收到 `ECONNREFUSED`。
- 标准安装器重试：`add . --list` 发现六个 Skill；临时项目的 Codex、Claude Code、Cursor 安装完成，安装后的各 `SKILL.md` 与源文件一致；Windows 默认模式安装 Codex 六个 Skill 完成。此前 npm 连接拒绝已不再阻断本地验证。
- 无 Go CLI、无 Git 的临时项目人工开发演练：初次测试因缺少模块失败；实现复用已有函数后，两项检查通过；写入 `PROJECT.md`、`STATUS.md`、`CHANGES.md`、`HANDOFF.md`，单独读取并复测通过。此项不属于真实 Agent 自动调用测试。
- 推送前重新运行 `.tools/go/bin/go.exe test ./... -count=1`、`vet ./...`、`build ./...`：均以退出码 0 完成。首次复检的临时目录位于本仓库内，使无 Git 场景的测试错误继承了仓库 Git 状态；改用系统临时目录复测后通过。Go 遥测令牌写入仍有权限提示，不影响退出码。
- 推送前 `git ls-remote origin refs/heads/master` 为 `5d8f22bff18aa76fad5271141c76331d3fecf26b`，与本地基线一致。
- 新版真实 Agent 自动触发、SkillHub 在线更新及 Codex/Claude Code/Cursor 跨客户端接手：未运行。

## 下一步与限制

用户已明确要求提交并推送本轮改动。推送后按 `docs/skills-first-plan.md` 的下一轮顺序验证公开 GitHub 安装、真实客户端自动调用与交接，再评估 SkillHub 更新。SkillHub 发布仍需另行明确要求。旧版 CLI 的 `checkpoint.json` 不应视为本轮有效检查点。
