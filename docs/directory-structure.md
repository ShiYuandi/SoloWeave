# 仓库目录说明

[English](en/directory-structure.md)

SoloWeave 是 AI 开发 Skills 项目。

| 路径 | 用途 |
| --- | --- |
| `skills/` | 五个 Skill 的唯一维护与安装来源。 |
| `skills/project-setup/references/` | 按需读取的选型判断指南，随 Skill 安装。 |
| `scripts/check-skills.ps1` | 检查发布元数据、引用和资源边界。 |
| `scripts/check-docs.ps1` | 检查 Markdown 非空及相对文件链接。 |
| `scripts/test-tooling.ps1` | 校验器负例、包内容与校验和回归检查。 |
| `scripts/prepare-behavior-tests.ps1`、`scripts/run-codex-behavior-tests.ps1` | 准备隔离项目并运行真实 Codex 会话；结果需人工核对。 |
| `scripts/package-skillhub.ps1` | 生成五个 ZIP 和 SHA256SUMS。 |
| `.github/workflows/ci.yml` | 运行维护检查并提供候选包。 |
| `README.md`、`CONTRIBUTING.md` | 项目入口与贡献约定。 |
| `docs/en/` | 公开说明的英文版。 |
| `docs/design.md`、`docs/development-plan.md` | 当前产品设计、实施与后续计划。 |
| `docs/configuration.md`、`docs/project-continuity.md` | 项目记录、决策和接手规则。 |
| `docs/agent-acceptance.md`、`docs/behavior-testing.md`、`docs/five-skills-acceptance.md`、`docs/end-to-end-acceptance.md` | 验收标准、回测方法与当前五项版本的实际结果；旧六项证据单独保留。 |
| `docs/codex-evidence/`、`docs/claude-code-evidence/` | 已执行回测的收据或原始证据。 |
| `docs/archive/` | 历史选型设计与已完成计划；当前规则以 docs/design.md 和 skills/ 为准。 |
| `docs/releasing.md`、`docs/release-status.md` | 发布步骤与已核对的线上状态。 |
| `.soloweave/context/`、`.soloweave/decisions/` | 本仓库的目标、进度、交接和架构决定。 |
| `AGENTS.md` | AI 在本仓库工作的约定。 |
| `LICENSE`、`SECURITY.md`、`CODE_OF_CONDUCT.md` | 许可证、安全报告和社区准则。 |
| `dist/skills/`、`dist/acceptance/` | 忽略的当前候选包与原始本地验收证据。 |
| `dist/archive/` | 可恢复的旧分发快照、测试样例和本机工具残留，不纳入 Git。 |
| `.tools/` | 可选、可重建的工具缓存；本轮已清理，后续生成仍会被忽略。 |
| `.git/` | Git 元数据。 |

其他项目可沿用已有文档，也可使用 `.soloweave/context/` 与 `.soloweave/decisions/` 保存上下文。安装和更新前先检查同名文件，保留用户修改。
