# 仓库目录说明

[English](en/directory-structure.md)

本页介绍 **SoloWeave 源码仓库**。CLI 安装到其他项目后生成的目录见[配置说明](configuration.md)和[项目连续性](project-continuity.md)。

## 根目录

| 路径 | 作用 | 是否纳入 Git |
| --- | --- | --- |
| `README.md` | 中文项目介绍、下载和快速开始 | 是 |
| `skills/` | 六个新版 Agent Skill 的唯一维护位置，可由标准安装器发现 | 是 |
| `AGENTS.md` | 仓库内 AI Agent 的工作约定 | 是 |
| `go.mod`、`go.sum` | Go 模块声明与依赖校验和 | 是 |
| `LICENSE` | MIT 许可证原文 | 是 |
| `CONTRIBUTING.md` | 贡献与验证约定 | 是 |
| `SECURITY.md` | 安全报告和敏感信息约定 | 是 |
| `CODE_OF_CONDUCT.md` | 社区行为准则 | 是 |
| `CHANGELOG.md` | 版本变更记录 | 是 |
| `.gitignore` | 排除本机工具链、构建产物和临时文件 | 是 |
| `.gitattributes` | 统一仓库文本文件换行，避免跨平台行尾差异 | 是 |
| `.git/` | Git 的本地元数据，由仓库所有者管理 | 否 |
| `.tools/` | 本机 Go 工具链、依赖缓存及验证器等，可再生 | 否 |
| `dist/` | Windows 下载包、本地预览程序及 `SHA256SUMS`，可再生 | 否 |

`.tools/` 和 `dist/` 保留在本机，方便继续开发与验收；它们不应进入提交。`dist/` 是构建输出，不代表已经在所有目标系统上运行过程序。

## 源码与构建

| 路径 | 作用 |
| --- | --- |
| `cmd/soloweave/main.go` | 命令行程序入口、退出码处理 |
| `internal/cli/` | Cobra 命令、交互输入和端到端 CLI 测试 |
| `internal/config/` | YAML 配置读写、JSON Schema 与 Go 语义校验；`schema/project.schema.json` 定义配置结构 |
| `internal/project/` | 项目初始化、架构批准、批准摘要和 ADR |
| `internal/continuity/` | 项目文档、Git 状态采集、检查点和恢复 |
| `internal/catalog/` | 可选技术栈预设目录的读取 |
| `internal/installer/` | Codex、Claude Code、Cursor 文件安装、预览、冲突与完整性检查 |
| `internal/bundle/` | 旧版 Go CLI 的嵌入资源；`assets/skills/` 是发布时的五个 Skill 快照，今后不作为新版源码编辑 |
| `scripts/package-skillhub.ps1` | 从根目录 `skills/` 打包六个新版 Skill 的本地候选 ZIP |
| `scripts/build-preview.ps1` | 构建四个目标平台的本地预览程序并生成校验和 |
| `scripts/package-windows.ps1` | 生成 Windows x64 ZIP、附带中英文说明并计算 SHA256 |
| `.github/workflows/ci.yml` | 远程仓库启用后执行 Go 测试、静态检查与构建 |
| `.github/workflows/release-windows.yml` | 手动运行 Windows 测试、打包并创建草稿 Release |

各 `*_test.go` 文件与同目录代码对应。修改嵌入资源后需重新构建 CLI，安装器才会使用新内容。

## 设计与交接

| 路径 | 作用 |
| --- | --- |
| `docs/development-plan.md` | V0.1 CLI 范围、步骤和验收历史 |
| `docs/design-v0.2.md` | 早期产品设计基准 |
| `docs/skills-first-design.md`、`docs/skills-first-plan.md` | 新版 Skills 优先设计与迁移计划 |
| `docs/configuration.md` | `project.yaml` 配置字段及批准流程 |
| `docs/project-continuity.md` | 项目交接文件和检查点说明 |
| `docs/skills.md` | 六个新版 Skill 的分工、安装与发布状态 |
| `docs/skill-only-workflow.md` | 无 CLI 时只用 Skills 的规划、检查与交接流程 |
| `docs/package-readme.zh-CN.md`、`docs/package-readme.en.md` | Windows ZIP 内双语说明的版本化模板 |
| `docs/agent-acceptance.md` | 尚待执行的真实客户端交接验收步骤 |
| `docs/directory-structure.md` | 本页 |
| `docs/en/` | README 和公开使用、贡献文档的英文版 |
| `.soloweave/project.yaml` | **本仓库自身**的已批准项目配置 |
| `.soloweave/approval.json` | 配置批准摘要；由 CLI 管理 |
| `.soloweave/decisions/ADR-0001.md` | 本仓库架构决策记录 |
| `.soloweave/context/PROJECT.md` | 本仓库的稳定目标与入口 |
| `.soloweave/context/STATUS.md` | 已完成、进行中、受阻与计划事项 |
| `.soloweave/context/HANDOFF.md` | 当前任务、实际验证、风险和下一步 |
| `.soloweave/context/CHANGES.md` | 重要改动记录 |
| `.soloweave/context/checkpoint.json` | 检查点的机器可读 Git 元数据；由 CLI 管理 |
| `.soloweave/context/tasks/` | 有任务细节时再放长期记录；空目录不会单独进入 Git |

`.soloweave/` 是项目知识的一部分，应随源码保留。检查点记录保存时的 Git 分支和项目文件内容；修改项目文件后，`context check` 会提示过期，需要新的交接检查点。只提交相同内容不会使新格式检查点过期。
