# 参与贡献

[English](docs/en/CONTRIBUTING.md)

较大改动前阅读 [项目设计](docs/design.md)、[开发计划](docs/development-plan.md)和[交接记录](.soloweave/context/HANDOFF.md)，再核对实际文件与 Git 状态。

普通用户直接安装 Skills。维护者运行仓库脚本需要 **PowerShell 7**，无需自行构建程序。

- 只在根目录 `skills/` 编辑 Skill。保持说明简洁，按需链接同一 Skill 内的参考资料。
- `SKILL.md` 中的 `name`、`description`、`license` 用于 Agent Skills；`slug`、`version`、`displayName`、`summary`、`homepage` 保留为 SkillHub 扩展元数据。修改内容时更新对应 Skill 的数字版本，不要求六个 Skill 同时升级。
- 重要行为变化用真实客户端场景验证；静态检查和安装成功不能代替实际调用。区分通过、失败、未运行和工具不可用。
- 完成重要修改后更新项目状态与交接；架构变化经开发者确认后写 ADR。
- 提交信息使用中文；提交、推送和发布须有开发者明确指令。

在仓库根目录运行：

```powershell
pwsh -NoProfile -File scripts/check-skills.ps1
pwsh -NoProfile -File scripts/check-docs.ps1
pwsh -NoProfile -File scripts/test-tooling.ps1
pwsh -NoProfile -File scripts/package-skillhub.ps1
```

校验器检查本仓库采用的单行元数据格式和资源边界，不是通用 YAML 解析器。工具测试覆盖无效元数据、缺失或越界资源、ZIP 内容、SHA256 及重复打包的一致性。输出在忽略的 `dist/skills/`。

发布步骤见[发布指南](docs/releasing.md)，客户端验证见[验收说明](docs/agent-acceptance.md)。不要提交工具缓存、构建产物、凭据或未审阅的本机会话日志。

按[行为回测指南](docs/behavior-testing.md)生成隔离夹具并运行已登录的 Codex CLI；此流程独立于不需要 Node.js 的静态 CI。
