# 五个 Skill 的独立安装与行为验收

日期：2026-10-11。当前源码版本见 [Skills 列表](skills.md)。安装文件检查与真实模型行为分别记录。

## 当前回归场景

每项使用新的隔离目录和会话，只安装该项，保留普通原始提示词、回复、工具调用和文件结果。

| Skill | 场景 | 应核对的行为 |
| --- | --- | --- |
| soloweave | 空目录中的模糊待办需求 | 缺少其他 Skill 仍澄清关键形态，待确认时不直接写代码。 |
| project-setup | 新全栈应用的选型 | 先理解需求，逐项比较候选并确认；参考指南不成为白名单。 |
| feature-workflow | 已有项目增加小功能 | 搜索复用已有代码，运行真实检查并留下必要记录；失败与未完成项如实记录。 |
| quality-review | 提供一组含具体缺陷的改动 | 按实际证据报告位置、影响和验证缺口，不虚报测试。 |
| project-continuity | 只有项目文件的新会话继续任务 | 核对目标、状态、文件与 Git，推进已知下一步，保持真实阻塞；不要求专门诊断根因。 |

本轮安装与回测结果见[验收记录](agent-acceptance.md)。不能要求所有普通任务调用全部五项。已有安装冲突合并不属于新目录安装核对的结论。

## 历史六项验收

移除故障排查前，skills@1.7.1 从当时源码发现六项并分别安装到六个临时项目。每次选择 Codex、Claude Code、Cursor 并使用复制模式；逐文件 SHA256 与当时源码一致，项目设置包含 references/selection-guide.md。Codex/Cursor 使用 .agents/skills/，Claude Code 使用 .claude/skills/，未写全局 Skills。

当时六项独立最小行为及额外接手场景共 12 个真实 Codex 会话通过，详见[历史行为结果](codex-behavior-results.md)。这些结果对应报告中的明确版本，不代替入口／功能 0.2.4 的本轮验证；首次读取权限阻塞见[历史尝试](codex-retest.md)。后续按[回测指南](behavior-testing.md)分别记录新客户端、模型和版本。
