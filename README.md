# SoloWeave

[English](docs/en/README.md)

**独立开发，稳妥交付。**

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 工程技能包与命令行工具。它把技术决策、任务进度和交接信息保存在项目文件中，让开发者切换账号或在 Codex、Claude Code、Cursor 之间切换时，能够从项目本身恢复上下文。

当前版本为 **V0.1 本地预览版**：提供 Go CLI、需要人工确认的 YAML 项目配置、五个 Agent Skill、平台规则安装和任务交接流程。不调用远程 AI API，也不需要数据库。[v0.2 设计基准](docs/design-v0.2.md)与[开发计划](docs/development-plan.md)记录了范围和取舍。

## 构建与测试

需要 Go 1.27.2 或更新版本。在仓库根目录运行：

```sh
go test ./...
go vet ./...
go build -o soloweave ./cmd/soloweave
```

Windows 可将输出文件名改为 `soloweave.exe`。`scripts/build-preview.ps1` 会在本地生成 Windows、Linux、macOS 预览程序和 SHA256 校验文件，输出在不纳入 Git 的 `dist/`。交叉编译成功不等于已在 Linux 或 macOS 上运行验证。

## 快速开始

先将构建出的 `soloweave` 程序加入 `PATH`，然后在目标项目目录中运行：

```sh
soloweave init
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave check
```

`init` 可交互选择预设或自定义技术栈；`soloweave catalog` 列出预设。自动化场景可使用 `init --preset api-go --name demo` 或 `init --from path/to/project.yaml`。两种方式都先生成**待确认**配置，仍须单独运行 `approve`。`approve --yes` 仅适合已经审阅配置的非交互场景。

任务暂停或切换 Agent 前，记录真实进度与验证结果：

```sh
soloweave context checkpoint --task "登录功能" --summary "表单已完成" --next "对接 API" --verification "go test ./...: 通过"
soloweave context resume
soloweave context check
```

没有运行测试时省略 `--verification`，交接记录会明确标为 `Not run`。没有 Git 仓库时，Git 状态会标为 `UNAVAILABLE`，不会假装已核验。详见[项目连续性](docs/project-continuity.md)和[配置说明](docs/configuration.md)。

## 文件放在哪里

目标项目的配置、架构决策记录和交接文件放在 `.soloweave/`。Codex 与 Cursor 共用 `.agents/skills/`，Claude Code 使用 `.claude/skills/`；平台规则分别写入 `AGENTS.md`、`CLAUDE.md` 与 `.cursor/rules/soloweave.mdc`。安装器预览改动并保护已有文件，发现冲突时停止覆盖。

本仓库各目录与文件的用途见[目录说明](docs/directory-structure.md)；Skill 和平台支持见[Skills 说明](docs/skills.md)。

## 验证范围与后续

Go 测试、静态检查、构建、五个 Skill 的格式校验，以及 Windows 预览程序在临时项目中的完整流程已通过。Linux、macOS 目前只完成交叉编译；Claude Code 与 Cursor 的真实客户端交接尚未实测，步骤见[跨 Agent 验收](docs/agent-acceptance.md)。远程 CI 与公开 Release 也尚未完成。

许可证：[MIT](LICENSE) · [参与贡献](CONTRIBUTING.md) · [安全问题](SECURITY.md) · [变更记录](CHANGELOG.md)
