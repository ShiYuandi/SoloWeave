# SoloWeave 项目

SoloWeave 帮助个人开发者和 1–5 人团队以已批准的工程决策、可复用工作流和项目文件中的交接信息使用 AI 编程 Agent。V0.1 提供 Go CLI、YAML 项目配置、五个 Agent Skill，以及 Codex、Claude Code、Cursor 的安装适配。

设计基准：`docs/design-v0.2.md`；开发计划：`docs/development-plan.md`；目录总览：`docs/directory-structure.md`。程序入口为 `cmd/soloweave/main.go`，核心包位于 `internal/`。使用 Go 1.27.2 运行 `go test ./...`、`go vet ./...` 和 `go build ./...`；`scripts/build-preview.ps1` 生成本地预览程序。仓库所有者负责远程发布。
