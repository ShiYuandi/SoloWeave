# 项目连续性

[English](en/project-continuity.md)

SoloWeave 把可供下一个 Agent 恢复的上下文放在 `.soloweave/context/`：

| 文件 | 用途 | 更新时机 |
| --- | --- | --- |
| `PROJECT.md` | 稳定的目标、模块、运行命令和入口 | 重要项目事实变化时 |
| `STATUS.md` | 已完成、进行中、受阻和待做事项 | 任务状态变化时 |
| `HANDOFF.md` | 当前任务、近期改动、实际验证、风险和下一步 | 保存检查点或交接时 |
| `CHANGES.md` | 已完成的重要修改 | 完成功能或重要修复时 |
| `tasks/` | 任务级详细记录 | 任务需要长期保存细节时 |

`context checkpoint` 写入 `HANDOFF.md` 和供 CLI 检查的 `checkpoint.json`。业务进度与实际验证结果由 Agent 或开发者提供；Git 可用时，CLI 采集分支、HEAD、工作区状态和项目文件内容摘要。`context check` 在分支或项目文件内容变化时提示过期；仅把相同内容提交到 Git 不会使新检查点过期。旧格式检查点仍使用原有 Git 状态判断，建议重新创建检查点。无 Git 时如实报告不可用。CLI 无法恢复检查点之后未记录的工作，也无法从代码自动推断业务意图。

新 Agent 应先运行 `context resume`，读取项目配置和任务记录，再检查代码及 Git 状态。如果实际文件与交接文档不一致，应按可能过期处理。不要在这些文件中写入 API Key、令牌、密码、私钥或真实密钥值。
