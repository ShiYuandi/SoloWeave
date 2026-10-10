# SoloWeave V0.1 本地开发计划

状态：V0.1 本地实现及自动化验收完成；GitHub Actions 已在 `b07d057` 提交通过。公开 Release 与真实 Claude Code/Cursor 客户端验收延后。基准：`docs/design-v0.2.md`（2026-10-09）。

## 目标与边界

交付面向个人开发者和 1–5 人团队的本地可用 V0.1：Go CLI、项目配置与架构确认、项目连续性、Codex/Claude Code/Cursor 安装器、五个核心 Agent Skills、检查与本地预览构建。SoloWeave 仓库由用户稍后自行创建；本次不执行 `git init`、推送或公开发布。Git 相关功能使用临时测试仓库验证。

开发顺序的第一项就是保存本计划及 v0.2 设计基准。此后任何实施取舍须更新本计划或 ADR，交接状态须记录在 `.soloweave/context/`。

## 阶段与验收

1. **工程基础（SW-001～006）**：本地模块名 `soloweave`，Cobra 命令骨架，MIT 许可证、README、Go 测试。准备工作区内不纳入版本控制的官方 Go 工具链。验收：`go build ./...`、`go test ./...`、帮助和版本命令可运行。
2. **配置与架构确认（SW-004～005、015～020）**：`.soloweave/project.yaml` 的 `schema_version: 1`；JSON Schema 验证形状，Go 校验跨字段语义。`init` 以交互或 `--from` 输入生成 draft；`approve` 显示配置、要求明确确认、生成 ADR 并转为 approved。支持 frontend/backend/fullstack，可自定义技术字符串；不将未知框架视为错误。验收：有效配置往返保存，错误配置给出字段定位；重大选型未经确认不生效。
3. **项目连续性（SW-007～014）**：初始化 `.soloweave/context/PROJECT.md`、`STATUS.md`、`HANDOFF.md`、`CHANGES.md`、`tasks/`；实现 `context show/checkpoint/resume/check`。`checkpoint` 从参数或交互读取任务摘要、下一步和真实验证结果，CLI 只自动采集 Git 状态，不推断业务进度。无 Git 时继续写检查点并明确标记 Git 不可用。验收：跨会话可由文档定位项目、当前任务、修改及下一步；缺失或过期信息可发现。
4. **Skills 与安装（SW-021～031）**：一个嵌入式资源源头，提供 `soloweave`、`project-setup`、`feature-workflow`、`project-continuity`、`quality-review`。Codex/Cursor 用 `.agents/skills/`，Claude Code 用 `.claude/skills/`；生成平台规则。`install --agents codex,claude,cursor --dry-run` 预览，正式安装记录 `.soloweave/installation.json`；重复安装无变化，用户文件冲突不覆盖，路径不逃逸项目。验收：三平台文件形状正确、安装可重复、冲突保留原文件。
5. **检查与本地交付（SW-032～038）**：`check` 检查项目契约、安装与上下文；`doctor` 报告环境及可用性；报告区分通过、失败、未执行、工具不可用。补齐 CI 配置、快速开始与人工跨 Agent 验收步骤，构建 Windows、Linux、macOS 本地预览包。公开仓库与 Release 待用户建立 Git 仓库后另行执行。

## 对外接口与数据约定

- CLI：`soloweave init [--from FILE]`、`approve`、`install [--agents LIST] [--dry-run]`、`doctor`、`check`、`version`、`context show|checkpoint|resume|check`。命令默认作用于当前目录；非交互使用必须能提供必要参数，不能暗自批准架构。
- `project.yaml` 是技术选型的唯一结构化事实来源；主要包含 `project`、按项目类型可选的 `frontend`/`backend`/`database`、`workspace`、`reuse`、`continuity`、`quality`、`agents`。路径仅允许项目内相对路径。批准后的关键字段通过再次明确批准与新 ADR 变更。
- `HANDOFF.md` 只保存当前任务、最近修改、实际验证、风险、下一步和 Git 状态。机器可验证的检查点元数据独立保存；源代码及 Git 状态优先于过期文档。敏感凭据不得写入上下文。
- 安装器仅管理自己创建的文件或带有 SoloWeave 标记的区块；记录内容摘要。已存在且不受管理的文件、或被用户修改过的受管理文件一律报告冲突，不静默替换。

## 验证清单

- `gofmt`、`go vet ./...`、`go test ./...`、`go build ./...`；交叉构建 Windows amd64、Linux amd64、macOS amd64/arm64。
- 配置测试：frontend/backend/fullstack、自定义技术、未知 schema 版本、无效组合、路径越界、draft/approved 与 ADR。
- 连续性测试：文件缺失、无 Git、干净与脏的临时 Git 仓库、检查点后代码变化、未运行的测试不得显示为通过。
- 安装测试：dry-run 不写入、首次安装、重复安装、用户文件冲突、安装状态损坏、路径与符号链接越界。
- 真实文件的端到端流程：初始化→批准→安装→任务检查点→无旧聊天记录的恢复。Claude Code/Cursor 当前未安装，以文件与临时仓库自动化验证；真实客户端验收步骤列于文档并标注未实测。

## 已确认的决定

- MIT，版权标识使用 `SoloWeave contributors`；本地 Go Module 暂为 `soloweave`，公开发布前再调整导入路径。
- 实施过程不执行 `git init`、提交或发布 Git 仓库。工作目录当前出现未提交的 Git 仓库，由用户自行管理；不依赖远程 AI API、数据库或 MCP Server。
- 技术实现采用 Go、Cobra、YAML、JSON Schema、Agent Skills、Git；Skill 内容按开放格式编写，详细规范按需加载。

## 本地交付记录

- 已实现配置、审批与 ADR、交接、五个内置 Skill、三个平台安装适配、诊断与检查命令。
- 自动化测试使用临时目录及临时 Git 仓库；四个平台的本地预览构建及 SHA256 校验文件位于忽略的 `dist/`。
- Claude Code 和 Cursor 的真实客户端交接仍待人工验收；步骤见 `docs/agent-acceptance.md`。远程地址现已配置，但 CI 工作流尚未推送和运行。

## 后续仓库整理（2026-10-09）

用户追加要求：将 README 等面向读者的文档改为中文，说明仓库目录，清理本机临时文件，完善 `.gitignore`，并明确授权创建本地 commit。此项覆盖本计划早期的“不提交”约束，但仍不授权推送、创建 Release 或宣称真实客户端验收已完成。

用户随后要求公开文档提供英文版，并选择暂不逐段翻译 v0.2 设计基准与历史开发计划。英文版位于 `docs/en/`，与中文页面互相链接；根目录 README 使用中文。首次提交的 Git 作者姓名和邮箱以用户明确提供的 `ShiYuandi`、`shiyuandi@foxmail.com` 为准，仅配置在本仓库。

## 检查点与模块路径后续修订（2026-10-10）

- 当前仓库远程地址已由所有者配置为 `https://github.com/ShiYuandi/SoloWeave.git`，Go Module 从初期临时名 `soloweave` 调整为 `github.com/ShiYuandi/SoloWeave`。
- 新检查点保存项目文件内容摘要及 Git 分支。项目内容变化仍提示过期；只提交相同内容不会造成误报。旧格式检查点仍按旧的 Git 状态规则检查，应重新创建以采用新规则。
- 只读核对 GitHub Actions：`b07d057` 和 `237d268` 两次 CI 均已完成并通过。后续改动的远程 CI 以对应提交的结果为准。
- 用户本轮要求暂不做跨平台实机与真实客户端验收；不推送或发布。

## Windows 免构建交付（2026-10-10）

- 用户确认普通用户应下载可直接运行的程序，并同意先提供 Windows x64 ZIP；跨平台实机验证继续延后。
- `scripts/package-windows.ps1` 从程序实际版本生成 ZIP，内含可执行文件、MIT 许可证及中英文 README，并生成 `SHA256SUMS`。本地解压后运行完整临时项目流程验收。
- `.github/workflows/release-windows.yml` 仅能由维护者在默认分支手动启动，运行测试、静态检查、打包与校验后创建草稿 Release；公开发布仍由仓库所有者检查并执行。
- 中英文 README 的普通用户入口改为 GitHub Releases 下载；源码构建移至贡献者部分。首个公开 Release 尚未发布；工作流只有进入默认分支后才能由维护者远程运行。
