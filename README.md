# SoloWeave

[English](docs/en/README.md)

**独立开发，稳妥交付。**

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 工程技能包与命令行工具。它把技术决策、任务进度和交接信息保存在项目文件中，让开发者切换账号或在 Codex、Claude Code、Cursor 之间切换时，能够从项目本身恢复上下文。

## 为什么我开发 SoloWeave

使用 AI 编程时，我发现它能很快写出一个功能，却不容易持续、规范地完成整个项目。从零开发全栈应用时，AI 可能直接替我选定技术栈和目录结构，后续又引入不同方案；它也可能忽略项目已有的组件和方法，写出重复的代码。

最让我困扰的是上下文丢失。切换账号、模型，或从 Codex 改用 Claude Code、Cursor 时，新的 AI 并不知道项目的目标、已经确认的架构、完成了什么、还有什么问题。我不得不反复介绍项目，有时还要让它重新分析代码。

我开始想：这些重要信息为什么只留在聊天记录里？如果技术决策、开发进度、代码变更、实际验证结果和下一步计划能保存在项目中，新的 AI 就可以先核对现状，再接着工作。

这就是我开发 SoloWeave 的原因。我希望 AI 在动手前和我讨论技术与架构选择，在开发时遵守已确认的约定、优先复用现有代码，并在重要节点留下清晰的交接记录。关键决定仍由开发者做，流程也不应让简单项目变得复杂。

SoloWeave 最初源于我自己的开发需求。我将它开源，希望其他独立开发者和小团队也能在不同 AI 工具之间延续工作：让 AI 帮我们持续推进项目，而不只是一次次生成代码。

## 核心能力

- **开发者确认架构**：先记录技术选择，再由开发者批准关键决策。
- **延续项目上下文**：把进度、变更、实际验证结果和下一步保存在项目中。
- **保持工程一致性**：让 Agent 先了解现有实现，并按已确认的约定开发与检查。

当前提供 **v0.1.0-preview Windows x64 预览版**：包含命令行程序、需要人工确认的 YAML 项目配置、五个 Agent Skill、平台规则安装和任务交接流程。不调用远程 AI API，也不需要数据库。[v0.2 设计基准](docs/design-v0.2.md)与[开发计划](docs/development-plan.md)记录了范围和取舍。

## 下载并使用 Windows 版

1. [下载 v0.1.0-preview Windows x64 安装包](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/soloweave-v0.1.0-preview-windows-amd64.zip)。其他版本见 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases)。
2. 解压 ZIP，得到 `soloweave.exe`。**普通用户无需安装 Go，也无需自行构建。**
3. 在 PowerShell 中运行下方的快速开始命令。

请下载发布页 **Assets** 中的 `soloweave-...-windows-amd64.zip`；GitHub 自动提供的 `Source code (zip)` 和 `Source code (tar.gz)` 是源码包，不含可直接运行的程序。安装包还包含中英文使用说明和 MIT 许可证；[SHA256SUMS](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/SHA256SUMS) 可用于核对下载文件。目前仅提供 Windows x64 安装包。

## 快速开始

将解压目录加入 `PATH`，然后在目标项目目录中运行。以下 PowerShell 示例中的路径请换成自己的解压目录和项目目录：

```powershell
$env:PATH = "C:\Tools\SoloWeave;$env:PATH"
cd C:\path\to\your-project
```

接着执行：

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

Go 测试、静态检查、构建、五个 Skill 的格式校验，以及 Windows 预览程序在临时项目中的完整流程已通过。Linux、macOS 目前只完成交叉编译；Claude Code 与 Cursor 的真实客户端交接尚未实测，步骤见[跨 Agent 验收](docs/agent-acceptance.md)。源码构建和参与开发的步骤见[贡献说明](CONTRIBUTING.md)。

许可证：[MIT](LICENSE) · [参与贡献](CONTRIBUTING.md) · [安全问题](SECURITY.md) · [变更记录](CHANGELOG.md)
