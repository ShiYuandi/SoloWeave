# 项目交接

核对日期：2026-10-11。

## 当前任务与结果

开发者已批准实施直到移除 Go 的八阶段计划，并要求继续补齐验收。本地迁移、六项独立最小行为与受阻跨会话接手已完成；随后明确授权的中文提交、推送、GitHub CI 与公开安装核对和六项 SkillHub 审核提交也已完成。Cursor、Anthropic 模型及更广场景另外记录；待审核与已发布分开记录。

先读 `docs/go-retirement-plan.md`、`docs/skills-first-design.md`、ADR-0003，再核对实际文件和 Git。旧 `docs/development-plan.md` 与 `docs/design-v0.2.md` 已标为历史。

## 工作区事实

- 分支 master；迁移提交 eb452f5、CI 修复 135e13c 已推送 origin/master。此后的交付记录提交以实际 Git 为准。
- 原有空目录路由修订、Claude Code 报告与证据已保留并随 Go 删除和维护脚本提交。
- `skills/` 是唯一当前 Skill 来源；入口/功能开发 0.2.3，设置 0.3.2，连续性 0.2.2，质量 0.2.1，调试 0.1.0。
- cmd/、internal/、go.mod、go.sum、Go 构建脚本和旧 CLI Release 工作流已删除；本地 Go 工具链与 CLI 包也已清理。
- 旧机器文件在 `docs/archive/go-cli/repository-state/`；当前没有机器批准或检查点。不要恢复旧文件来描述新工程，也不要删除其他项目的资料。
- `dist/skills/` 有六个本地候选 ZIP 与 SHA256SUMS，受 .gitignore 忽略。维护者需 PowerShell 7，用户只需安装 Skills。
- 既有 CLI Release 保留，未创建新 Release。SkillHub 已更新五项并通过 GitHub 导入新增故障排查；六项审核均已提交，见 docs/release-status.md。

## 实际验证与边界

1. check-skills.ps1 通过：六项单行元数据、唯一 slug、版本与本地资源边界；不是通用 YAML 校验。
2. test-tooling.ps1 的 15 项检查通过：负例拒绝、包输出边界、六包校验和与全部资源字节、重复打包一致。
3. package-skillhub.ps1 成功；check-docs.ps1 相对文件链接与 git diff --check 通过。网页链接和标题锚点不在脚本范围；修复负例残留退出码后 135e13c 远端 CI 通过。
4. 标准安装器 skills@1.7.1 在六个全新临时项目分别复制一个 Skill，并选择 Codex/Claude Code/Cursor；安装资源与源码 SHA256 一致。未测试覆盖现有用户修改。
5. 三个归档机器文件与 HEAD 原内容一致（换行归一化）。原返回 ZIP 目前已不在先前路径，不能重复声称本轮又做过 ZIP 字节比对。
6. 导入的 Claude Code CLI 原始证据来自前一组版本、实际模型 glm-5.3；场景一看到 soloweave → project-setup → feature-workflow，场景二看到 soloweave → project-setup，受阻场景只看到 soloweave。本机未重跑应用。
7. 首次 Codex CLI 读取拒绝保留为历史；后续使用当前候选版、默认配置模型 gpt-6.1-sol，保留沙箱和执行规则，完成 12 个真实临时会话。六项独立行为、全套新项目路由、缺失工具双会话、获准修复通过。每项都有实际读取本地 Skill 的事件；独立断言、原测试和构建指令哈希、失败/修复结果经外部核对。详见 docs/codex-behavior-results.md，公开收据在 docs/codex-evidence/2026-10-11.json，原始日志在 dist/acceptance/codex-retirement-followup/。
8. quick_validate.py 缺 PyYAML 未成功。用户全局 Skills 未重装。
9. 标准安装器 skills@1.7.1 从公开 GitHub 成功安装六项到全新临时项目，选择 Codex/Claude Code/Cursor；两套平台目录共 14 文件与源码逐项 SHA256 一致。安装目标记录在 dist/acceptance/public-install-path.txt。

## 继续执行

- 本轮最小验收已完成。按 docs/behavior-testing.md 可用准备与运行脚本复跑；夹具拒绝覆盖已有目录，证据拒绝覆盖同名文件，结果需人工核对，退出码 0 不能自动算通过。
- 不能把 Claude Code 等同于 Anthropic 模型。Cursor、Anthropic 模型及更广场景尚未完成，不从 Codex 的通过推断其他客户端相同。
- 原始会话已检查常见密钥模式且未发现匹配，机器路径等原始资料保留；已按明确授权执行提交与线上步骤。后续提交仍使用中文。
- 不要因客户端权限问题无依据改写路由，不要自动全局安装。后续新增发布工作仍需遵守 AGENTS.md 的授权规则。

## 本轮线上操作授权

开发者明确授权的操作已执行。GitHub 源码可安装，135e13c CI 通过；SkillHub 项目交接 0.2.2 核对时已发布，其余五项安全审核中。后续先核对平台结果，再验收公开下载，不能将提交成功写成全部发布完成。截图在 dist/acceptance/skillhub-submitted-2026-10-11.jpg，详细快照见 docs/release-status.md。
