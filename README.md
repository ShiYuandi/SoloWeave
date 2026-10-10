# SoloWeave

[English](docs/en/README.md)

**独立开发，稳妥交付。**

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 工程技能包与命令行工具。它把技术决策、任务进度和交接信息保存在项目文件中，让开发者切换账号或在 Codex、Claude Code、Cursor 之间切换时，能够从项目本身恢复上下文。

## 为什么我开发 SoloWeave

使用 AI 编程时，我发现它**能很快写出一个功能，却不容易持续、规范地完成整个项目**。从零开发全栈应用时，AI 可能**未经我确认就选定技术栈和目录结构**，后续又引入不同方案；它也可能忽略项目已有的组件和方法，**重复编写相似功能**。

最让我困扰的是**开发上下文丢失**。切换账号、模型，或从 Codex 改用 Claude Code、Cursor 时，新的 AI 并不知道项目的目标、已经确认的架构、完成了什么、还有什么问题。我不得不**反复介绍项目，让 AI 重新分析已有工作**。

我开始想：这些重要信息为什么只留在聊天记录里？如果能**把技术决策、开发进度、代码变更、实际验证结果和下一步计划保存在项目本身**，新的 AI 就可以先核对现状，再接着工作。

这就是我开发 SoloWeave 的原因。我希望 AI **在动手前和我讨论技术与架构选择**，在开发时**遵守已确认的约定、优先复用现有代码**，并在重要节点**留下清晰的交接记录**。关键决定仍由开发者做，流程也不应让简单项目变得复杂。

SoloWeave 最初源于我自己的开发需求。我将它开源，希望其他独立开发者和小团队也能在不同 AI 工具之间延续工作：**让 AI 帮我们持续推进项目，而不只是一次次生成代码。**

## 核心能力

- **开发者确认架构**：先记录技术选择，再由开发者批准关键决策。
- **延续项目上下文**：把进度、变更、实际验证结果和下一步保存在项目中。
- **保持工程一致性**：让 Agent 先了解现有实现，并按已确认的约定开发与检查。

当前提供 **v0.1.0-preview Windows x64 预览版**：包含命令行程序、需要人工确认的 YAML 项目配置、五个 Agent Skill、平台规则安装和任务交接流程。不调用远程 AI API，也不需要数据库。[v0.2 设计基准](docs/design-v0.2.md)与[开发计划](docs/development-plan.md)记录了范围和取舍。

## 只装 Skills，还是使用 CLI？

| 使用方式 | 你会得到什么 | 当前限制 |
| --- | --- | --- |
| **只安装五个 Skills** | AI 会按规范讨论架构、检查并复用现有代码，还能在项目文档中手工记录经你确认的决策、进度和交接。只想使用这些开发规范时，**不必安装 CLI**。 | 不会生成 CLI 的批准记录、安装记录和可自动检查的交接点；`soloweave check` 等 CLI 命令不可用。AI 应如实报告这些检查未运行。 |
| **使用 CLI（由 CLI 安装 Skills）** | 在上述规范之外，用命令创建项目配置和交接文件、记录批准的架构决策、安装五个 Skills 与项目规则，并检查文件状态。 | 目前只提供 Windows x64 的可运行 CLI 包。 |

CLI 是本地文件工具，不会替 AI 写业务代码，也不会替开发者批准技术方案或凭空判断测试通过。

## 让 AI 安装 Skills

如果你的 AI 编程工具可以访问网络并安装 Skill，可以把下面的提示词直接发送给它：

```text
请根据 https://skillhub.cn/install/skillhub.md，为当前 AI 编程工具安装以下五个 SoloWeave Skills：
@user_38c0807e/shiyuandi-soloweave
@user_38c0807e/shiyuandi-soloweave-project-setup
@user_38c0807e/shiyuandi-soloweave-feature-workflow
@user_38c0807e/shiyuandi-soloweave-project-continuity
@user_38c0807e/shiyuandi-soloweave-quality-review
安装前检查已有的同名文件，不要覆盖我修改过的内容。安装后核对五个 SKILL.md，并告诉我安装路径和版本。
```

入口 Skill 需要其余四个配合。这段提示词只安装 Skill 文件，**只用开发规范可以到此为止**。安装后，可以对 AI 说“请按 SoloWeave 的纯 Skills 流程开发，先检查现有实现；手工记录经我确认的决策和交接，说明未执行的 CLI 检查”。具体文件和验收方式见[纯 Skills 使用指南](docs/skill-only-workflow.md)。

如果选择下方的 CLI 路线，**不必先从 SkillHub 安装**：`soloweave install` 本身会安装五个 Skills 和平台规则。SkillHub 安装的文件不会被 CLI 的 `check`/`doctor` 视为 CLI 管理的项目安装；若两种方式已写入同一目录，先预览并处理同名文件冲突。详见 [Skills 说明](docs/skills.md)。

## 下载并使用 Windows 版

1. [下载 v0.1.0-preview Windows x64 安装包](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/soloweave-v0.1.0-preview-windows-amd64.zip)。其他版本见 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases)。
2. 解压 ZIP，得到 `soloweave.exe`。**普通用户无需安装 Go，也无需自行构建。**
3. 在 PowerShell 中运行下方的快速开始命令。

请下载发布页 **Assets** 中的 `soloweave-...-windows-amd64.zip`；GitHub 自动提供的 `Source code (zip)` 和 `Source code (tar.gz)` 是源码包，不含可直接运行的程序。安装包还包含中英文使用说明和 MIT 许可证；[SHA256SUMS](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/SHA256SUMS) 可用于核对下载文件。目前仅提供 Windows x64 安装包。

## CLI 命令做什么

在目标项目目录运行这些命令；它们操作的是该项目中的文件。

| 命令 | 作用 |
| --- | --- |
| `soloweave catalog` | 查看可选的技术栈预设。 |
| `soloweave init` | 创建**待确认**的 `.soloweave/project.yaml`，并初始化项目与交接文档；不会替你批准技术方案。 |
| `soloweave approve` | 展示当前配置，征得开发者确认后标记为已批准，并生成架构决策记录（ADR）。 |
| `soloweave install --dry-run` | 预览将写入的五个 Skills 和平台规则，不修改文件。正式 `install` 才写入，并保护已有或被修改的文件。 |
| `soloweave context checkpoint` | 根据你或 AI 提供的任务摘要、下一步和**实际**验证结果写入交接记录；可用时采集 Git 状态，不自动推断进度或测试是否通过。 |
| `soloweave context resume` / `context show` | 读取项目和交接资料，帮助新会话接手。 |
| `soloweave context check` | 检查交接资料是否缺失或过期。 |
| `soloweave check` | 检查项目配置、架构批准、CLI 安装记录和交接状态。 |
| `soloweave doctor` | 报告 Git、项目配置和 CLI 安装状态，帮助定位环境问题。 |
| `soloweave version` | 显示 CLI 版本。 |

## 快速开始

将解压目录加入 `PATH`，然后在目标项目目录中运行。以下 PowerShell 示例中的路径请换成自己的解压目录和项目目录：

```powershell
$env:PATH = "C:\Tools\SoloWeave;$env:PATH"
cd C:\path\to\your-project
```

接着执行以下 **CLI 路线（包含 Skills）**。`--agents codex,claude,cursor` 表示为这三个平台安装项目文件；只用其中一部分时，改成需要的平台即可。

```sh
soloweave init
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave context checkpoint --task "初始化项目" --summary "已确认配置并安装 Skills" --next "开始开发"
soloweave check
```

`init` 可交互选择预设或自定义技术栈；`soloweave catalog` 列出预设。自动化场景可使用 `init --preset api-go --name demo` 或 `init --from path/to/project.yaml`。两种方式都先生成**待确认**配置，仍须单独运行 `approve`。`approve --yes` 仅适合已经审阅配置的非交互场景。

任务暂停或切换 Agent 前，记录真实进度与验证结果：

```sh
soloweave context checkpoint --task "登录功能" --summary "表单已完成" --next "对接 API"
soloweave context resume
soloweave context check
```

只有实际运行了验证，才在 `context checkpoint` 中添加 `--verification "命令: 实际结果"`。

没有运行测试时省略 `--verification`，交接记录会明确标为 `Not run`。没有 Git 仓库时，Git 状态会标为 `UNAVAILABLE`，不会假装已核验。详见[项目连续性](docs/project-continuity.md)和[配置说明](docs/configuration.md)。

## 文件放在哪里

目标项目的配置、架构决策记录和交接文件放在 `.soloweave/`。Codex 与 Cursor 共用 `.agents/skills/`，Claude Code 使用 `.claude/skills/`；平台规则分别写入 `AGENTS.md`、`CLAUDE.md` 与 `.cursor/rules/soloweave.mdc`。安装器预览改动并保护已有文件，发现冲突时停止覆盖。

本仓库各目录与文件的用途见[目录说明](docs/directory-structure.md)；Skill 和平台支持见[Skills 说明](docs/skills.md)。

## 验证范围与后续

Go 测试、静态检查、构建、五个 Skill 的格式校验，以及 Windows 预览程序在临时项目中的完整流程已通过。Linux、macOS 目前只完成交叉编译；Claude Code 与 Cursor 的真实客户端交接尚未实测，步骤见[跨 Agent 验收](docs/agent-acceptance.md)。源码构建和参与开发的步骤见[贡献说明](CONTRIBUTING.md)。

许可证：[MIT](LICENSE) · [参与贡献](CONTRIBUTING.md) · [安全问题](SECURITY.md) · [变更记录](CHANGELOG.md)
