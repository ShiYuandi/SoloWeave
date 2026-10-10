# 项目交接

核对日期：2026-10-10。

## 当前任务

将 SoloWeave 从 CLI 为主迁移为 Skills 优先、AI 主动使用的开发工作流。本轮完成项目选型规则的本地修订，下一阶段是按验收场景在真实新会话复测推荐与交接行为。

## 当前事实

分支 `master`，Skills 优先基线已以 `5bc8f1e` 推送；三个 Skill 的 `0.2.1` 修订、`project-setup` `0.3.0`、项目选型设计与计划、双语文档和交接记录尚未提交，以 Git 状态为准。最新公开的 GitHub Release 仍为 `v0.1.1-preview` 旧版 CLI；SkillHub 仍是五个旧版 Skill，尚无新版发布。本仓库 `project.yaml` 和 `approval.json` 仍代表旧版 Go 子工程；整体新方向见 `ADR-0002.md`。

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
- 用户提供的 Codex 双窗口试用：六个全局安装 `SKILL.md` 与根目录源文件 SHA256 一致；仓库外 `test/` 项目只有 `PROJECT.md` 与 ADR，缺少 `STATUS.md`、`HANDOFF.md`、`CHANGES.md`。窗口 1 明确报告本机 pnpm 阻止原生构建脚本；窗口 2 对“请继续这个项目”称应用可运行并要求另选方向。此样本的跨会话交接未通过，内部 Skill 自动调用状态不可由摘要确定。详情见 `docs/agent-acceptance.md`。
- 用户确认受阻交接与接手行为的有限修正；已修改三个仓库 Skill，版本均为 `0.2.1`，并更新中英文验收与版本说明。`pwsh -NoProfile -File scripts/package-skillhub.ps1` 通过；六个当前 ZIP 的条目与源文件内容、SHA256SUMS 核对通过；`git diff --check` 通过。尚未用修订版在真实 Codex 客户端复测。
- 用户指出 `project-setup` 的技术推荐仅有少量固定组合。经设计与实施计划确认，已修订 `skills/project-setup/SKILL.md` 至 `0.3.0`，新增 `references/selection-guide.md`，并修改 `scripts/package-skillhub.ps1` 打包全部所需文件。类型用于发现决策维度，参考示例不限制 AI 选型。同步了中英文使用、验收及版本文档；真实 Agent 新会话行为尚未验收。
- 本轮运行 `pwsh -NoProfile -File scripts/package-skillhub.ps1` 成功；核对六个 ZIP 的 SHA256SUMS、条目和逐文件字节一致，其中 `project-setup` 有两个文件。标准安装器 `skills@1.7.1 add . --list` 找到六个 Skill；隔离项目复制安装的 `project-setup` 两个文件与源码 SHA256 一致。改动 Markdown 相对链接检查通过。首次普通 `npx` 因命令不存在、`pnpm dlx` 因 npm 连接拒绝失败；后来经授权的标准安装器运行成功。`quick_validate.py` 缺少 PyYAML，未完成。
- 新版 Skill 的真实自动调用事件不可由现有对话摘要确认；修订后的 `0.2.1` 尚未在客户端复测。SkillHub 在线更新与 Claude Code/Cursor 真实客户端接手未运行。

## 下一步与限制

用户已明确要求将本轮修改提交并推送到 GitHub，结果以 Git 状态及远端为准；尚未要求更新 SkillHub。之后将 `project-setup` `0.3.0` 安装到真实新会话，按双语验收文档复测需求澄清、维度推荐、表外技术和确认边界；另需将三个 `0.2.1` Skill 安装到测试客户端复测交接，并核对是否真实触发。用户全局安装的 Skill 仍是先前试用版。旧版 CLI 的 `checkpoint.json` 不应视为本轮有效检查点。
