# 参与贡献

[English](docs/en/CONTRIBUTING.md) · 中文为主版本

修改较大功能前，请阅读[开发计划](docs/development-plan.md)、[v0.2 设计基准](docs/design-v0.2.md)和[当前交接记录](.soloweave/context/HANDOFF.md)，再核对实际代码与 Git 状态。

- Skill 尽量简洁、与平台无关；平台差异放在安装适配和规则文件中。
- 修改 CLI 行为时，补充针对行为的测试；不要用只复述实现的测试替代真实验证。
- 重要改动及时更新 `.soloweave/context/STATUS.md`、`CHANGES.md`，交接时更新 `HANDOFF.md`。
- 有意改变已批准的关键架构时，重新审阅项目配置，执行 `soloweave approve` 并保留新的 ADR。
- 提交前运行 `go test ./...`、`go vet ./...`、`go build ./...`，如实说明未运行的检查。

提交中不要包含 `.tools/`、`dist/`、密钥或其他本机生成文件。目录用途见[目录说明](docs/directory-structure.md)。
