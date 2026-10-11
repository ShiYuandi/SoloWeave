# 项目记录与技术决策

[English](en/configuration.md)

当前 SoloWeave 使用 Skills 和 Markdown 项目记录，不需要初始化命令、YAML 配置或专用程序。

优先沿用项目已有的 README、开发计划、状态和决策文档。需要建立持续交接记录时，默认使用：

| 文件 | 内容 |
| --- | --- |
| `.soloweave/context/PROJECT.md` | 项目目标、主要模块、运行入口及已确认选型。 |
| `.soloweave/context/STATUS.md` | 当前任务、已完成、阻塞和下一步。 |
| `.soloweave/context/HANDOFF.md` | 最近修改、真实验证、风险及接手步骤。 |
| `.soloweave/context/CHANGES.md` | 重要修改与涉及文件。 |
| `.soloweave/decisions/ADR-NNNN.md` | 关键选择、候选、理由、约束和开发者确认事实。 |

候选技术和架构在开发者确认前保持“待确认”；确认后由 Agent 记录，不需要 `approve` 命令。后续重大变化另建 ADR，说明取代的决定。类型和参考指南不限制可选择的技术。

Skills 不提供机器批准摘要或自动过期检测。接手时核对记录、实际文件和 Git 状态；不能仅凭文档里的“已完成”判断项目可运行。详见[项目连续性](project-continuity.md)。

已有旧 CLI 项目请按[迁移指南](cli-migration.md)处理；[旧 YAML 字段说明](archive/go-cli/configuration.md)仅供历史参考。
