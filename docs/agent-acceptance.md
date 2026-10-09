# 跨 Agent 交接验收

[English](en/agent-acceptance.md) · 中文为主版本

当前开发机未安装 Claude Code 或 Cursor。自动化测试已覆盖安装产物和检查点行为，但这两个客户端中的真实交接**尚未验证**。

安装客户端后，按以下步骤进行人工验收：

1. 新建临时项目，依次运行 `soloweave init`、`soloweave approve`、`soloweave install --agents codex,claude,cursor`。
2. 在 Agent A 中开始一个小功能，按实际进度更新 `STATUS.md`、`CHANGES.md`。运行真实检查，保存包含准确结果的检查点。
3. 关闭 Agent A；在没有旧聊天记录的情况下，用 Agent B 打开同一项目，只让它读取项目文件继续工作。
4. 确认 Agent B 能找到项目目标、已批准技术栈、完成和剩余工作、近期改动、真实验证结果及下一步。检查点之后改动一个源文件，确认它能发现交接内容可能过期。
5. 让 Agent B 完成任务、运行检查并更新状态与交接文档；确认它没有绕过开发者确认和 ADR 修改关键架构。

执行时记录客户端版本及结果。在完成前，不应宣称已验证真实 Claude Code/Cursor 客户端兼容性。
