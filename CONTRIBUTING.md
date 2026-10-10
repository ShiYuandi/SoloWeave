# 参与贡献

[English](docs/en/CONTRIBUTING.md)

修改较大功能前，请阅读[Skills 优先设计](docs/skills-first-design.md)、[迁移计划](docs/skills-first-plan.md)和[当前交接记录](.soloweave/context/HANDOFF.md)，再核对实际代码与 Git 状态。早期 CLI 范围保留在[开发计划](docs/development-plan.md)和[v0.2 设计基准](docs/design-v0.2.md)。

修改根目录 `skills/` 中的新版 Skill 不需要 Go。只有维护旧版 CLI 才需要 Go 1.27.2 或更新版本；普通用户可直接安装 Skills，不必下载程序或自行构建。

- 新版 Skill 只在根目录 `skills/` 编辑；`internal/bundle/assets/skills/` 是旧版 CLI 快照。Skill 尽量简洁、与平台无关。
- 修改 CLI 行为时，补充针对行为的测试；不要用只复述实现的测试替代真实验证。
- 重要改动及时更新 `.soloweave/context/STATUS.md`、`CHANGES.md`，交接时更新 `HANDOFF.md`。
- 有意改变已批准的关键架构时，先取得开发者明确确认并新增 ADR；旧版 CLI 自身配置变更仍按其 `approve` 流程处理。
- 修改新版 Skill 时，运行 `scripts/package-skillhub.ps1` 并核对 ZIP；修改旧版 CLI 时再运行 `go test ./...`、`go vet ./...`、`go build ./...`。如实说明未运行的检查。
- Git commit 的提交信息使用中文，简要说明本次改动。

提交中不要包含 `.tools/`、`dist/`、密钥或其他本机生成文件。目录用途见[目录说明](docs/directory-structure.md)。

## 从源码构建

以下命令仅用于维护旧版 Go CLI，在仓库根目录运行：

```sh
go test ./...
go vet ./...
go build -o soloweave.exe ./cmd/soloweave
```

## Windows 下载包

在 Windows PowerShell 中运行 `./scripts/package-windows.ps1`，会在忽略的 `dist/` 下生成 Windows x64 ZIP 和 `SHA256SUMS`。版本号直接取自包内 `soloweave.exe version`，避免包名与程序版本不一致。发布前请解压 ZIP、运行其中的程序并核对校验值。

代码提交并推送到默认分支后，仓库所有者可在 GitHub Actions 手动运行 **Windows draft release**。工作流会运行 Go 测试与静态检查，打包并验证程序，然后创建带有 ZIP 和校验文件的**草稿 Release**。核对草稿内容后，由仓库所有者在 GitHub 上发布；运行工作流本身不会公开发布草稿。
