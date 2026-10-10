# SoloWeave

[English](docs/en/README.md)

**独立开发，稳妥交付。**

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 开发 Skills。安装后，你可以像平常一样提出开发需求；AI 会按任务读取项目事实、遵守已确认的决策、检查可复用代码，并在重要节点维护项目记录。

## 为什么我开发 SoloWeave

使用 AI 编程时，我发现它**能很快写出一个功能，却不容易持续、规范地完成整个项目**。从零开发全栈应用时，AI 可能**未经我确认就选定技术栈和目录结构**，后续又引入不同方案；它也可能忽略项目已有的组件和方法，**重复编写相似功能**。

最让我困扰的是**开发上下文丢失**。切换账号、模型，或从 Codex 改用 Claude Code、Cursor 时，新的 AI 并不知道项目的目标、已经确认的架构、完成了什么、还有什么问题。我不得不**反复介绍项目，让 AI 重新分析已有工作**。

我开始想：这些重要信息为什么只留在聊天记录里？如果能**把技术决策、开发进度、代码变更、实际验证结果和下一步计划保存在项目本身**，新的 AI 就可以先核对现状，再接着工作。

这就是我开发 SoloWeave 的原因。我希望 AI **在动手前和我讨论技术与架构选择**，在开发时**遵守已确认的约定、优先复用现有代码**，并在重要节点**留下清晰的交接记录**。关键决定仍由开发者做，流程也不应让简单项目变得复杂。

SoloWeave 最初源于我自己的开发需求。我将它开源，希望其他独立开发者和小团队也能在不同 AI 工具之间延续工作：**让 AI 帮我们持续推进项目，而不只是一次次生成代码。**

## 快速开始

仓库根目录的 [`skills/`](skills/) 包含六个可独立安装的 Skill。可以让支持 [Agent Skills](https://agentskills.io/specification) 的 AI 工具从本仓库安装，也可以使用 [Skills CLI](https://github.com/vercel-labs/skills) 选择安装：

```sh
npx skills@latest add ShiYuandi/SoloWeave
```

也可以把下面这段发给你的 AI，让它处理安装：

```text
请根据 https://github.com/vercel-labs/skills 的安装说明，从 https://github.com/ShiYuandi/SoloWeave 仓库的 skills/ 目录安装六个 SoloWeave Skills 到当前 AI 编程工具。先检查同名文件，不覆盖我的修改；安装后报告各 SKILL.md 的路径和版本。以后我正常提出开发需求时，请按适用的 SoloWeave Skill 工作。
```

安装后直接提出普通需求，例如：

> 帮我给这个项目增加用户登录。先看已有实现，完成后运行相关测试，并留下下一次能接手的进度记录。

无需每次输入 `/setup`、`/build` 或 `soloweave` 命令。入口 Skill 会按任务选择项目设置、功能开发、故障排查、质量检查或交接流程。首次进入长期开发的项目时，AI 会检查现有文档和代码；重大技术选型仍需你确认。自动调用由各 AI 工具决定，不能保证每次触发；必要时可明确说“使用 SoloWeave 继续这个项目”。

**当前验证状态：**标准安装器已从本地仓库发现六个新版 Skill，并在隔离项目生成 Codex、Claude Code 和 Cursor 对应的 Skill 文件；真实客户端自动调用、公开 GitHub 地址安装与 SkillHub 更新仍待逐项验收。现有 [SkillHub 页面](docs/skills.md#skillhub-发布状态)提供的是此前发布的五个旧版 Skill，更新前请按页面版本识别。

## 它会做什么

| 时机 | AI 应做的事 |
| --- | --- |
| 进入已有项目 | 读取规则、目标、决策、状态和相关代码；核对交接是否过期。 |
| 创建或接入项目 | 分析现状，说明必要选项；重大技术与架构选择由你确认，再写入 ADR。 |
| 开发或修复 | 先查找可复用实现，再改代码并运行相关验证。 |
| 完成重要任务或切换账号 | 更新项目状态、重要变更、真实测试结果和下一步，让新会话从文件接手。 |

项目记录优先沿用已有文档；需要新建时放在 `.soloweave/context/` 和 `.soloweave/decisions/`。只装 Skills **不需要下载 Go CLI**。纯 Skills 路线不生成 CLI 的机器批准、安装或检查点文件，也不声称运行了 CLI 检查。详见[使用指南](docs/skill-only-workflow.md)和[六个 Skill 的分工](docs/skills.md)。

## 旧版可选 CLI

此前发布的 [v0.1.1-preview Windows x64 包](https://github.com/ShiYuandi/SoloWeave/releases/tag/v0.1.1-preview)仍可下载，无需自行构建。它提供 `init`（待确认配置）、`approve`（人工确认并记录 ADR）、`install`（安装当时内置的五个旧版 Skill）、`context checkpoint/resume/check`（交接）、`check`（项目校验）和 `doctor`（环境诊断）。该程序是发布时的快照，**不会自动获得根目录 `skills/` 中的新版流程**。已有用户可继续使用；新项目建议先使用上方 Skills 路线。CLI 的完整命令说明见[旧版配置文档](docs/configuration.md)与[项目连续性说明](docs/project-continuity.md)。

| 命令 | 作用 |
| --- | --- |
| `catalog` | 查看旧版技术栈预设。 |
| `init` / `approve` | 创建待确认配置；开发者审阅后批准并生成 ADR。 |
| `install --dry-run` / `install` | 预览或安装内置旧版 Skills 与平台规则。 |
| `context checkpoint` / `resume` / `check` | 保存任务摘要与真实验证、读取交接、检查资料是否过期。 |
| `check` / `doctor` / `version` | 校验项目文件、诊断环境、查看程序版本。 |

## 验证与参与

新版 Skills 的本地格式、打包和无 CLI 流程验收结果会记录在[验收文档](docs/agent-acceptance.md)；真实客户端结果以该页为准。[目录说明](docs/directory-structure.md)介绍仓库文件，[贡献说明](CONTRIBUTING.md)介绍源码开发。

许可证：[MIT](LICENSE) · [参与贡献](CONTRIBUTING.md) · [安全问题](SECURITY.md) · [变更记录](CHANGELOG.md)
