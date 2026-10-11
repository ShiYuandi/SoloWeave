# 仓库目录说明

[English](en/directory-structure.md)

SoloWeave 是 AI 开发 Skills 项目。

| 路径 | 用途 |
| --- | --- |
| `skills/` | 六个 Skill 的唯一维护与安装来源。 |
| `skills/project-setup/references/` | 按需读取的选型判断指南，随 Skill 安装。 |
| `scripts/check-skills.ps1` | 检查发布元数据、引用和资源边界。 |
| `scripts/check-docs.ps1` | 检查 Markdown 非空及相对文件链接。 |
| `scripts/test-tooling.ps1` | 校验器负例、包内容与校验和回归检查。 |
| `scripts/prepare-behavior-tests.ps1`、`scripts/run-codex-behavior-tests.ps1` | 准备隔离项目并运行真实 Codex 会话；结果需人工核对。 |
| `scripts/package-skillhub.ps1` | 生成六个 ZIP 和 SHA256SUMS。 |
| `.github/workflows/ci.yml` | 运行维护检查并提供候选包。 |
| `README.md`、`CONTRIBUTING.md` | 项目入口与贡献约定。 |
| `docs/en/` | 公开说明的英文版。 |
| `docs/design.md`、`docs/development-plan.md` | 当前产品设计、实施与后续计划。 |
| `docs/configuration.md`、`docs/project-continuity.md` | 项目记录、决策和接手规则。 |
| `docs/agent-acceptance.md`、`docs/behavior-testing.md` | 验收标准、实际结果与回测方法。 |
| `docs/codex-evidence/`、`docs/claude-code-evidence/` | 已执行回测的收据或原始证据。 |
| `docs/superpowers/` | 项目选型流程的详细设计与实施记录。 |
| `docs/releasing.md`、`docs/release-status.md` | 发布步骤与已核对的线上状态。 |
| `.soloweave/context/`、`.soloweave/decisions/` | 本仓库的目标、进度、交接和架构决定。 |
| `AGENTS.md` | AI 在本仓库工作的约定。 |
| `LICENSE`、`SECURITY.md`、`CODE_OF_CONDUCT.md` | 许可证、安全报告和社区准则。 |
| `.tools/`、`dist/` | 忽略的工具缓存、候选包与本地验证日志。 |
| `.git/` | Git 元数据。 |

其他项目可沿用已有文档，也可使用 `.soloweave/context/` 与 `.soloweave/decisions/` 保存上下文。安装和更新前先检查同名文件，保留用户修改。
