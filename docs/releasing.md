# 发布 Skills

[English](en/releasing.md)

当前交付物是六个 Skill ZIP 和 SHA256SUMS。普通用户从 GitHub 安装即可，不需要维护者发布可执行程序。

1. 审阅实际改动和客户端证据，为内容变化的 Skill 更新数字 `X.Y.Z` 版本；保留既有 slug。同步中英文文档与 CHANGELOG。
2. 运行 `scripts/check-skills.ps1`、`scripts/check-docs.ps1`、`scripts/test-tooling.ps1`、`scripts/package-skillhub.ps1`，核对 `dist/skills/SHA256SUMS` 列出的六个 ZIP 和参考文件。
3. 开发者明确要求后，以中文提交并推送。确认 GitHub CI 对该提交通过；远程未运行时不写“CI 已通过”。
4. 从公开地址在隔离项目安装，核对发现的六个名称、版本、参考文件和内容。当前未推送的本地修订不会出现在 GitHub 安装中。
5. 开发者明确要求发布时，在 GitHub 创建新版本说明；可附加六个 ZIP 和 SHA256SUMS。Release 版本代表整套快照，每个 Skill 的版本仍独立。
6. 核对并更新六个 SkillHub slug。优先核对 GitHub 导入预览；也可上传对应 ZIP。逐项确认版本、中文概述、正文和参考文件，再提交审核。
7. 审核通过后核对公开页面与实际下载，更新仓库发布状态。提交审核、审核通过和公开可下载是不同状态。

GitHub Actions 的 Skills CI 会保存候选包 artifact；artifact 不是公开 Release。SkillHub 页面仍为旧版时，应如实标明，不能把本地打包当成线上更新。

安装器参数以 [vercel-labs/skills 官方说明](https://github.com/vercel-labs/skills)为准；本仓库不提供新的安装器。更新已有同名文件前先检查冲突，保留用户修改。
