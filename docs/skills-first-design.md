# SoloWeave Skills 优先设计

日期：2026-10-10。状态：实施基准。

## 目标

开发者安装一次 SoloWeave Skills 后，直接提出普通开发需求。Agent 根据任务主动读取项目事实、选择适用工作流、复用已有代码、完成验证，并在有意义的检查点维护项目记录。切换账号或 Agent 时，新会话只依赖仓库文件即可接手。开发者仍决定关键技术与架构选型。

## 产品边界

- 默认使用路径只有 Agent Skills 和项目内的 Markdown 文件，不依赖 Go 程序、远程服务或特定聊天记录。
- 已发布的 Go CLI 继续可用，但作为旧版可选工具维护；新版 Skills 不调用它，也不把 CLI 的机器检查视为已运行。
- Skills 可以被 Agent 自动选择，但格式规范不保证每次调用。项目规则可提高发现概率，仍需在真实客户端验收。手动点名入口 Skill 是补救方式，不是日常前提。
- 不自动批准架构，不覆盖现有项目规则或文档，不替用户执行提交、推送、发布等需单独授权的操作。

## 仓库结构与兼容

`skills/` 是今后唯一编辑的 Skill 源码目录，包含 `soloweave` 入口及 `project-setup`、`feature-workflow`、`systematic-debugging`、`quality-review`、`project-continuity`。标准安装器从根目录发现它们。`internal/bundle/assets/skills/` 暂保留为 Go CLI v0.1.1-preview 的旧版内置快照，不同步新版内容；后续在纯 Skills 验收完成后再决定移除 Go。发布脚本从根目录 `skills/` 打包。

五个既有 Skill 保持名称与 SkillHub slug，以便逐项更新。新增调试 Skill 在实际发布前只作为 GitHub 仓库中的新 Skill；不能把本地打包成功描述为 SkillHub 已上架。所有面向读者的中文文档为主要入口，英文版覆盖 README 和使用说明。

## Agent 工作流

1. **进入项目**：入口 Skill 判断是新项目还是已有项目。已有项目先读 README、相关代码、`.soloweave/context/` 和 ADR；若存在已批准的 `project.yaml`，将其作为结构化约束。核对 Git 与源文件；资料缺失或过期时说明差异。
2. **选择任务流程**：新项目或重要选型用 `project-setup`；功能开发用 `feature-workflow`；故障修复用 `systematic-debugging`；审查与交付检查用 `quality-review`；恢复和检查点用 `project-continuity`。入口只负责路由和共同约束，各 Skill 保持可独立使用。
3. **实施**：先搜索相关实现和测试，判断是否复用；重大选型先向开发者说明取舍并取得明确确认。普通实现、检查和文档更新按任务授权继续进行。
4. **验证与交接**：运行适用的项目命令，分别记录通过、失败、未运行和工具不可用。完成有意义的任务、出现阻塞或准备切换会话时，更新 `STATUS.md`、`CHANGES.md`、`HANDOFF.md`。小型文字或格式修改只需适度记录，不强制生成任务文档。
5. **恢复**：新 Agent 先读交接资料，再核对当前 Git、代码和测试；发现不一致时以实际文件为准，不把旧记录当成事实。

## 文件契约

- `.soloweave/context/PROJECT.md`：稳定目标、运行入口及已确认决策链接。
- `.soloweave/context/STATUS.md`：当前任务、已完成事项、阻塞与下一步。
- `.soloweave/context/CHANGES.md`：重要改动及对应文件。
- `.soloweave/context/HANDOFF.md`：当前任务、最近修改、真实验证结果、风险、下一步、核对时间及可取得的 Git 状态。
- `.soloweave/decisions/ADR-NNNN.md`：开发者明确确认后的重要决定与理由。待确认内容不得写成已批准。

优先沿用目标项目已有的等价文档。纯 Skills 流程不创建 `approval.json`、`checkpoint.json` 或 `installation.json`；它们属于旧版 CLI 的机器记录。涉及凭据时只记配置要求，不写实际密钥。

## 验收标准

- 标准安装器可发现根目录的全部六个 Skill；SKILL.md 符合 Agent Skills 规范。
- 不提供 Go CLI 的干净项目中，Agent 可完成项目分析、经确认的决策记录、开发或修复、真实验证和交接。
- 无旧聊天记录的新会话能定位目标、当前任务、最近改动、验证结果与下一步，并识别过期记录。
- 普通开发请求的自动选择在可用客户端实测；未实测的平台明确标注。自动调用不能宣称百分之百保证。
- Go CLI 旧版测试保持通过，现有用户的安装文件和数据不会被迁移脚本覆盖。
