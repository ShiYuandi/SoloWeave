# Codex CLI 回测记录

> 下文记录首次受阻尝试，保留历史。本轮后续已完成六项独立最小场景与跨会话检查，见[最新结果](codex-behavior-results.md)；不能继续把旧阻塞当作当前唯一结论。

日期：2026-10-11。客户端：codex-cli 0.162.0-alpha.17.2。

在仓库外临时目录复制六个 Skills 到 .agents/skills/，发送原始空目录待办提示词。测试时版本为 soloweave/feature-workflow 0.2.2、project-setup 0.3.1、project-continuity 0.2.1、quality-review 0.2.0、systematic-debugging 0.1.0。

首次启动因本机 Codex 会话目录权限失败，退出码 1。允许写会话目录后，客户端实际启动并退出 0，但读取项目目录及 Skill 的命令被执行策略拒绝；回复报告无法读取或创建文件。虽然启动参数为 workspace-write，实际会话报告只读限制。

结论：**行为验收受阻，不能算通过，也不能归因于 Skill 路由失败**。没有生成待办应用。原始本地事件与错误输出保存在忽略的 `dist/acceptance/codex-retirement/`，没有更改全局 Skills。退出码 0 只表示 CLI 结束，不代表验收成功。

该尝试之后已在保留工作区沙箱和执行规则的真实新会话重测，当前候选版本的独立最小场景、空目录路由、受阻接手与获准修复检查通过。原始日志在 dist/acceptance/codex-retirement-followup/。Cursor、Anthropic 模型及更多场景仍需独立验证；不将这次启动成功归因于未经单独验证的某一个参数。
