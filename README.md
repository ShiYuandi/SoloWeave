# SoloWeave

[English](docs/en/README.md)

**独立开发，稳妥交付。**

SoloWeave 是面向个人开发者和 1–5 人团队的 AI 开发 Skills。安装后，你可以像平常一样提出开发需求；AI 会按任务读取项目事实、遵守已确认的决策、检查可复用代码，并在重要节点维护项目记录。

## 为什么我开发 SoloWeave

使用 AI 编程时，我发现它**能很快写出一个功能，却不容易持续、规范地完成整个项目**。

从零开发项目时，AI 通常会提出技术方案并让我确认，但我仍可能不清楚**为什么推荐这些技术、还有哪些适合的选择，以及不同方案的取舍**。我希望它能**先了解需求，再按前端、后端、数据库和架构等适用维度逐项讨论**，让我在理解方案后作出决定。

随着开发推进，AI 还可能偏离已经确认的工程约定，或忽略项目已有的组件和方法，**重复编写相似功能**。

最让我困扰的是**开发上下文丢失**。切换账号、模型，或从 Codex 改用 Claude Code、Cursor 时，新的 AI 并不知道项目的目标、已经确认的架构、完成了什么、还有什么问题。我不得不**反复介绍项目，让 AI 重新分析已有工作**。

我开始想：这些重要信息为什么只留在聊天记录里？如果能**把技术决策、开发进度、代码变更、实际验证结果和下一步计划保存在项目本身**，新的 AI 就可以先核对现状，再接着工作。

这就是我开发 SoloWeave 的原因。我希望 AI **在动手前和我讨论技术与架构选择**，在开发时**遵守已确认的约定、优先复用现有代码**，并在重要节点**留下清晰的交接记录**。关键决定仍由开发者做，流程也不应让简单项目变得复杂。

SoloWeave 最初源于我自己的开发需求。我将它开源，希望其他独立开发者和小团队也能在不同 AI 工具之间延续工作：**让 AI 帮我们持续推进项目，而不只是一次次生成代码。**

## 六个 Skill：做什么，什么时候介入

每个 Skill 负责一个工作环节，**可以整套安装，也可以按需独立安装**。下表说明各项适用的时机；AI 应根据当前任务选择，不需要每次把六项全部运行一遍。

| Skill | 干什么用 | 什么时候介入 |
| --- | --- | --- |
| **开发入口** [soloweave](skills/soloweave/SKILL.md) | 先读目录、规则和项目记录，判断新旧项目，再选择适用流程。 | 开始开发、规划、修复、审查，或说“继续这个项目”时。 |
| **项目设置** [project-setup](skills/project-setup/SKILL.md) | 澄清需求，按适用维度比较技术与架构方案；关键选择由你确认，再记录决定与理由。 | 空目录创建项目、为已有项目建立工程约定，或需要重要技术选型时。接入已有项目先尊重现有实现。 |
| **功能开发** [feature-workflow](skills/feature-workflow/SKILL.md) | 查找可复用代码，遵守已确认的约定，实施功能、运行相关检查并更新进度。 | 在已有项目中新增功能或修改代码时；新项目先完成必要的需求与选型确认。 |
| **故障排查** [systematic-debugging](skills/systematic-debugging/SKILL.md) | 先复现问题、收集证据和定位原因，再做适度修复并回归验证。 | 出现报错、测试失败、行为异常或性能回退时。 |
| **质量检查** [quality-review](skills/quality-review/SKILL.md) | 核对实际改动、工程约定和验证结果，报告具体问题、位置与剩余缺口。 | 请求代码审查、交付前检查，或确认项目是否能交接时。 |
| **项目交接** [project-continuity](skills/project-continuity/SKILL.md) | 保存真实进度、修改、验证和下一步；接手时核对文件，再推进已知工作。 | 开始新会话、完成重要任务、遇到阻塞，或准备切换账号、模型与 AI 工具时。 |

### 它们怎样协作

- **“从零做一个待办应用。”** 开发入口识别新项目 → 项目设置澄清需求、讨论选型 → 你确认关键方案后，功能开发实施 → 在重要节点记录交接。
- **“这个接口报错，帮我修好。”** 开发入口核对现状 → 故障排查复现、定位和修复 → 验证后记录结果；需要审查时再进行质量检查。
- **“换了一个账号，请继续这个项目。”** 项目交接读取目标、状态和下一步，核对实际代码与验证结果，再按任务进入功能开发或故障排查。仍有阻塞就如实记录，不把未完成工作当成已完成。

自动选择由 AI 客户端决定，不能保证每次触发；必要时可以明确点名，例如“使用 project-setup 帮我讨论选型”。更多说明见[Skills 列表](docs/skills.md)。

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

安装后正常描述开发需求即可。入口 Skill 会按任务选择项目设置、功能开发、故障排查、质量检查或交接流程。首次进入长期开发的项目时，AI 会检查现有文档和代码；重大技术选型仍需你确认。自动调用由各 AI 工具决定，不能保证每次触发；必要时可明确说“使用 SoloWeave 继续这个项目”。

**当前验证状态：**六项独立安装文件核对、静态检查与打包通过；当前源码版本也通过了真实 Codex CLI 的六项独立最小场景、空目录路由及受阻接手回测。12 个会话的证据和边界见[行为结果](docs/codex-behavior-results.md)与[验收记录](docs/agent-acceptance.md)。此前 Claude Code CLI 的证据使用较早版本，实际模型为 `glm-5.3`；Cursor 和 Anthropic 模型本轮未实测。GitHub 与 SkillHub 的实际更新进度见[交付状态](docs/release-status.md)。

## 项目记录留在哪里

项目记录优先沿用已有文档；需要新建时放在 `.soloweave/context/` 和 `.soloweave/decisions/`。**安装 Skills 就可以使用，无需下载或构建程序。** Skills 维护 Markdown 记录，实际验证使用目标项目已有命令。详见[使用指南](docs/skill-only-workflow.md)和[六个 Skill 的分工](docs/skills.md)。

## 验证与参与

本地格式、打包和真实客户端结果见[验收文档](docs/agent-acceptance.md)。[目录说明](docs/directory-structure.md)介绍仓库文件，[贡献说明](CONTRIBUTING.md)介绍维护者检查，[发布指南](docs/releasing.md)说明 GitHub 与 SkillHub 更新流程。

许可证：[MIT](LICENSE) · [参与贡献](CONTRIBUTING.md) · [安全问题](SECURITY.md) · [变更记录](CHANGELOG.md)
