# 用 SoloWeave Skills 开发

[English](en/skill-only-workflow.md)

安装[六个新版 Skill](skills.md)后，直接向 AI 描述开发任务。常规工作不需要先运行 SoloWeave CLI，也不需要在每次对话中输入固定斜杠命令。AI 应自动检查项目已有规则、源代码和交接资料，再选择适合的工作流。自动选择由各客户端实现，若未触发，可明确点名 `soloweave`。

## 初次进入项目

AI 先查看 README、依赖文件、相关代码与测试。已有 `.soloweave/` 时，再核对 `PROJECT.md`、`STATUS.md`、`HANDOFF.md` 和 ADR 与当前文件、Git 是否一致。没有 SoloWeave 文档的现有项目继续沿用已有规范；计划长期开发或需要跨会话接力时，再建立最少的上下文文件。重大架构决定由开发者确认后记入 ADR，待确认方案不得标成批准。

## 正常开发与交接

功能开发先搜索可复用实现；故障修复先取得可观察的复现证据。运行项目实际可用的测试或检查，并准确区分通过、失败、未运行和工具不可用。完成重要任务、受阻或准备切换账号时，AI 应维护 `.soloweave/context/` 下的 `STATUS.md`、`CHANGES.md` 和简短 `HANDOFF.md`。新的 AI 先读记录，再核对 Git 和源文件；有冲突时以实际文件为准。

`PROJECT.md` 保留稳定目标和已确认决策链接；`HANDOFF.md` 至少能找到当前任务、最近改动、实际验证与结果、遗留问题、下一步和核对时间。Git 不可用就明确写不可用。不要写入密钥。纯 Skills 不生成 Go CLI 的 `approval.json`、`checkpoint.json` 或 `installation.json`，也不把人工核对称为 `soloweave check` 通过。

详细验收步骤见[跨 Agent 验收](agent-acceptance.md)。
