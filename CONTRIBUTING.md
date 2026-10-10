# 参与贡献

[English](docs/en/CONTRIBUTING.md)

修改较大功能前，请阅读[开发计划](docs/development-plan.md)、[v0.2 设计基准](docs/design-v0.2.md)和[当前交接记录](.soloweave/context/HANDOFF.md)，再核对实际代码与 Git 状态。

- Skill 尽量简洁、与平台无关；平台差异放在安装适配和规则文件中。
- 修改 CLI 行为时，补充针对行为的测试；不要用只复述实现的测试替代真实验证。
- 重要改动及时更新 `.soloweave/context/STATUS.md`、`CHANGES.md`，交接时更新 `HANDOFF.md`。
- 有意改变已批准的关键架构时，重新审阅项目配置，执行 `soloweave approve` 并保留新的 ADR。
- 提交前运行 `go test ./...`、`go vet ./...`、`go build ./...`，如实说明未运行的检查。
- Git commit 的提交信息使用中文，简要说明本次改动。

提交中不要包含 `.tools/`、`dist/`、密钥或其他本机生成文件。目录用途见[目录说明](docs/directory-structure.md)。

## Windows 下载包

在 Windows PowerShell 中运行 `./scripts/package-windows.ps1`，会在忽略的 `dist/` 下生成 Windows x64 ZIP 和 `SHA256SUMS`。版本号直接取自包内 `soloweave.exe version`，避免包名与程序版本不一致。发布前请解压 ZIP、运行其中的程序并核对校验值。

代码提交并推送到默认分支后，仓库所有者可在 GitHub Actions 手动运行 **Windows draft release**。工作流会运行 Go 测试与静态检查，打包并验证程序，然后创建带有 ZIP 和校验文件的**草稿 Release**。核对草稿内容后，由仓库所有者在 GitHub 上发布；运行工作流本身不会公开发布草稿。
