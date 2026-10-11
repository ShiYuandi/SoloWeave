# 项目交接

核对日期：2026-10-11。

## 当前任务与结果

开发者已批准实施直到移除 Go 的八阶段计划，并要求继续补齐验收。本地迁移、六项独立最小行为与受阻跨会话接手已完成；下一步按明确指令提交和发布。Cursor、Anthropic 模型及更广场景另外记录；不要把本地检查通过写成在线发布完成。

先读 `docs/go-retirement-plan.md`、`docs/skills-first-design.md`、ADR-0003，再核对实际文件和 Git。旧 `docs/development-plan.md` 与 `docs/design-v0.2.md` 已标为历史。

## 工作区事实

- 分支 master，HEAD 仍为 fd6e04a77e32bdff0d121fda753be4231a6f1333。本轮未提交、推送或修改远端。
- 原有未提交的空目录路由修订、Claude Code 报告与证据被保留；当前 Git 差异还包含 Go 删除和新维护脚本。
- `skills/` 是唯一当前 Skill 来源；入口/功能开发 0.2.3，设置 0.3.2，连续性 0.2.2，质量 0.2.1，调试 0.1.0。
- cmd/、internal/、go.mod、go.sum、Go 构建脚本和旧 CLI Release 工作流已删除；本地 Go 工具链与 CLI 包也已清理。
- 旧机器文件在 `docs/archive/go-cli/repository-state/`；当前没有机器批准或检查点。不要恢复旧文件来描述新工程，也不要删除其他项目的资料。
- `dist/skills/` 有六个本地候选 ZIP 与 SHA256SUMS，受 .gitignore 忽略。维护者需 PowerShell 7，用户只需安装 Skills。
- 上次记录的公开 Release 为旧 CLI v0.1.1-preview，SkillHub 为五个旧版 Skill；本轮未重新在线核对或更新。

## 实际验证与边界

1. check-skills.ps1 通过：六项单行元数据、唯一 slug、版本与本地资源边界；不是通用 YAML 校验。
2. test-tooling.ps1 的 15 项检查通过：负例拒绝、包输出边界、六包校验和与全部资源字节、重复打包一致。
3. package-skillhub.ps1 成功；check-docs.ps1 相对文件链接与 git diff --check 通过。网页链接、标题锚点及远端 CI 未验证。
4. 标准安装器 skills@1.7.1 在六个全新临时项目分别复制一个 Skill，并选择 Codex/Claude Code/Cursor；安装资源与源码 SHA256 一致。未测试覆盖现有用户修改。
5. 三个归档机器文件与 HEAD 原内容一致（换行归一化）。原返回 ZIP 目前已不在先前路径，不能重复声称本轮又做过 ZIP 字节比对。
6. 导入的 Claude Code CLI 原始证据来自前一组版本、实际模型 glm-5.3；场景一看到 soloweave → project-setup → feature-workflow，场景二看到 soloweave → project-setup，受阻场景只看到 soloweave。本机未重跑应用。
7. 首次 Codex CLI 读取拒绝保留为历史；后续使用当前候选版、默认配置模型 gpt-6.1-sol，保留沙箱和执行规则，完成 12 个真实临时会话。六项独立行为、全套新项目路由、缺失工具双会话、获准修复通过。每项都有实际读取本地 Skill 的事件；独立断言、原测试和构建指令哈希、失败/修复结果经外部核对。详见 docs/codex-behavior-results.md，公开收据在 docs/codex-evidence/2026-10-11.json，原始日志在 dist/acceptance/codex-retirement-followup/。
8. quick_validate.py 缺 PyYAML 未成功。用户全局 Skills 未重装。

## 继续执行

- 本轮最小验收已完成。按 docs/behavior-testing.md 可用准备与运行脚本复跑；夹具拒绝覆盖已有目录，证据拒绝覆盖同名文件，结果需人工核对，退出码 0 不能自动算通过。
- 不能把 Claude Code 等同于 Anthropic 模型。Cursor、Anthropic 模型及更广场景尚未完成，不从 Codex 的通过推断其他客户端相同。
- 提交前审阅 docs/claude-code-evidence/ 的原始会话及路径信息，按 docs/releasing.md 执行明确授权的发布步骤。提交信息中文。
- 不要因客户端权限问题无依据改写路由，不要自动全局安装，不要自行提交、推送或更新 SkillHub。

## 本轮线上操作授权

开发者已明确授权中文提交、推送，确认 GitHub CI 与公开安装后提交六项 SkillHub 审核。当前正在执行，实际进度见 docs/release-status.md；不把待审核写成已发布。
