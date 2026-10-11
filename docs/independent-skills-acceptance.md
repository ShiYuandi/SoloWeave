# 六个 Skill 的独立安装与行为验收

日期：2026-10-11。当前版本见 [Skills 列表](skills.md)。

## 本轮已执行

本地缓存的标准安装器 skills@1.7.1 从当前源码发现六个 Skill。在六个分别新建的临时项目中，每次只选一个 Skill，明确选择 Codex、Claude Code、Cursor 并使用复制模式；逐文件核对安装后的内容与源码 SHA256 一致。项目设置包含 references/selection-guide.md。

该安装器版本为 Codex/Cursor 写入 .agents/skills/，为 Claude Code 写入 .claude/skills/；整个测试不写全局 Skills。此前 pnpm dlx 包装命令在本机报“skills 不是命令”，随后直接运行已缓存的官方 CLI JavaScript 成功。

以上是安装文件验证。随后已在真实 Codex CLI 用当前候选版本完成六项独立最小场景，实际读取本地 Skill 文件；功能复用、故障修复、审查定位、待确认选型和跨会话接手均通过。12 个会话和外部核对详见[行为结果](codex-behavior-results.md)。首次权限阻塞作为[历史记录](codex-retest.md)保留。

## 回归场景与继续扩展

每项使用新的隔离目录和会话，只安装该项，保留普通原始提示词、回复、工具调用和文件结果。

| Skill | 场景 | 应核对的行为 |
| --- | --- | --- |
| soloweave | 空目录中的模糊待办需求 | 缺少其他 Skill 仍澄清关键形态，待确认时不直接写代码。 |
| project-setup | 新全栈应用的选型 | 先理解需求，逐项比较候选并确认；参考指南不成为白名单。 |
| feature-workflow | 已有项目增加小功能 | 搜索复用已有代码，运行真实检查并留下必要记录；空目录先处理项目设置。 |
| systematic-debugging | 有失败测试的最小项目 | 先复现并定位，修复后重跑；不删掉失败检查。 |
| quality-review | 提供一组含具体缺陷的改动 | 按实际证据报告位置、影响和验证缺口，不虚报测试。 |
| project-continuity | 只有项目文件的新会话继续任务 | 核对目标、状态、文件与 Git，推进已知下一步，保持真实阻塞。 |

本轮上述最小场景在 Codex 通过；这不保证所有任务自动触发。后续按[回测指南](behavior-testing.md)复跑，并对新客户端和模型分别记录结果；不能要求所有普通任务调用六个 Skill。更新现有安装时先检查用户修改，已有安装冲突合并不属于上述新目录安装测试的结论。
