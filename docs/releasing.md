# 发布 Skills

[English](en/releasing.md)

当前交付物是五个 Skill ZIP 和 SHA256SUMS。普通用户从 GitHub 安装即可，不需要维护者发布可执行程序。

1. 审阅实际改动和客户端证据，为内容变化的 Skill 更新数字 `X.Y.Z` 版本；保留既有 slug。同步中英文文档与 CHANGELOG。
2. 运行 `scripts/check-skills.ps1`、`scripts/check-docs.ps1`、`scripts/test-tooling.ps1`、`scripts/package-skillhub.ps1`，核对 `dist/skills/SHA256SUMS` 列出的五个 ZIP 和参考文件；发布以清单为准，不混入输出目录可能保留的旧版本包。
3. 开发者明确要求后，以中文提交并推送。确认 GitHub CI 对该提交通过；远程未运行时不写“CI 已通过”。
4. 从公开地址在隔离项目安装，核对发现的五个名称、版本、参考文件和内容。当前未推送的本地修订不会出现在 GitHub 安装中。
5. 开发者明确要求发布时，在 GitHub 创建新版本说明；可附加五个 ZIP 和 SHA256SUMS。Release 版本代表整套快照，每个 Skill 的版本仍独立。
6. 核对并更新五个 SkillHub slug。优先核对 GitHub 导入预览；也可上传对应 ZIP。逐项确认版本、中文概述、正文和参考文件，再提交审核。
7. 审核通过后核对公开页面与实际下载，更新仓库发布状态。提交审核、审核通过和公开可下载是不同状态。

GitHub Actions 的 Skills CI 会保存候选包 artifact；artifact 不是公开 Release。SkillHub 页面仍为旧版时，应如实标明，不能把本地打包当成线上更新。

安装器参数以 [vercel-labs/skills 官方说明](https://github.com/vercel-labs/skills)为准；本仓库不提供新的安装器。更新已有同名文件前先检查冲突，保留用户修改。

## Agensi 免费发布

1. 取得开发者对该平台发布的明确授权。每个 ZIP 包含单层 Skill 目录、SKILL.md 及所需参考文件；现有打包脚本可生成该结构。逐资源比对 skills/，记录 ZIP 哈希。
2. 登录创作者后台，先核对同名条目避免重复。每项上传一个包，选择 **Free / No charge**。本项目的平台名称和介绍使用英文；上传包保留中文源码，并在介绍注明语言和真实源码版本。
3. 审阅或清空平台自动填写的内容。删除不存在的命令、未经验证的兼容性承诺和虚构演示；填写实际文件／终端权限及条件性联网需求、MIT 来源链接和验证边界。
4. 预览免费价格、正文和资源树后提交。分别记录提交、审核、公开可见和真实 Agent 调用结果；平台可能显示独立包标记 v1.0，不能替代正文中的源码版本。
5. 审核后核对公开英文页面、Free 标记与参考资料，更新[交付状态](release-status.md)。公开页面检查不等于 MCP 加载内容已实测。

格式及流程参考：[官方包结构清单](https://www.agensi.io/learn/skill-md-creator-checklist)、[官方发布指南](https://www.agensi.io/learn/how-to-sell-skills-on-agensi)。文档与表单不一致时以实际页面为准；2026-10-18 免费提交未要求连接收款账户或支付费用。
