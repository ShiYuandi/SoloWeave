# 仓库目录说明

[English](en/directory-structure.md)

SoloWeave 当前是纯 Skills 项目，源码中不再包含 Go CLI。

| 路径 | 用途 |
| --- | --- |
| `skills/` | 六个 Skill 的唯一维护与安装来源，每个目录至少含一个 SKILL.md。 |
| `skills/project-setup/references/` | 按需读取的选型判断指南，随项目设置 Skill 一起安装。 |
| `scripts/check-skills.ps1` | 检查名称、发布元数据、引用和资源边界。 |
| `scripts/check-docs.ps1` | 检查当前与历史 Markdown 的相对文件链接。 |
| `scripts/test-tooling.ps1` | 校验器负例、打包内容与校验和的回归测试。 |
| `scripts/prepare-behavior-tests.ps1`、`scripts/run-codex-behavior-tests.ps1` | 生成隔离项目并运行真实 Codex 会话；结果需人工核对。 |
| `scripts/package-skillhub.ps1` | 生成六个 ZIP 和 SHA256SUMS；同一内容重复打包结果一致。 |
| `.github/workflows/ci.yml` | 在 GitHub Actions 运行同一套检查，并提供候选包下载。 |
| `README.md`、`CONTRIBUTING.md` | 中文项目入口和贡献约定。 |
| `docs/en/` | 公开使用、贡献和发布文档的英文版。 |
| `docs/configuration.md`、`docs/project-continuity.md` | Markdown 项目记录、决策和接手规则。 |
| `docs/agent-acceptance.md` | 实际客户端证据与尚未完成的验证。 |
| `docs/cli-migration.md` | 旧 Go CLI 停更及项目迁移说明。 |
| `docs/releasing.md` | GitHub 与 SkillHub 的发布流程。 |
| `docs/archive/go-cli/` | 旧 CLI 使用说明及本仓库旧机器记录；仅供历史参考。 |
| `docs/design-v0.2.md`、`docs/development-plan.md` | 早期 CLI 设计和开发历史。 |
| `.soloweave/context/` | 本仓库的目标、状态、变更和当前交接。 |
| `.soloweave/decisions/` | 架构决策；ADR-0003 确认退出 Go。 |
| `AGENTS.md` | AI 在本仓库工作的约定。 |
| `LICENSE`、`SECURITY.md`、`CODE_OF_CONDUCT.md` | MIT 许可证、安全报告和社区行为准则。 |
| `.tools/`、`dist/` | 忽略的本机工具缓存、候选包和本地验证证据；不纳入提交。 |
| `.git/` | Git 元数据，由仓库所有者管理。 |

当前目录没有 `cmd/`、`internal/`、`go.mod` 或 Go 构建工作流。旧版源码可从历史提交或 Release 标签找回。其他项目使用 SoloWeave 后可沿用已有文档，或创建 `.soloweave/context/` 与 `.soloweave/decisions/`；不会强制生成程序配置。
