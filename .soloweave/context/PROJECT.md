# SoloWeave 项目

SoloWeave 帮助个人开发者和 1–5 人团队以已批准的工程决策、可复用工作流和项目文件中的交接信息使用 AI 编程 Agent。新版以根目录 `skills/` 的六个 Agent Skill 为主体，开发者正常提出任务，AI 主动选择流程并维护上下文。已发布的 Go CLI 暂作为旧版可选工具保留。

当前方向与迁移步骤：`docs/skills-first-design.md`、`docs/skills-first-plan.md`；批准依据：`ADR-0002.md`。V0.1 的历史基准仍见 `docs/design-v0.2.md`、`docs/development-plan.md`。目录总览：`docs/directory-structure.md`。旧版 CLI 入口为 `cmd/soloweave/main.go`，核心包位于 `internal/`；本仓库的 `.soloweave/project.yaml` 与 `approval.json` 仍描述该旧版 Go 子工程，不代表新版 Skills 需要 Go。仓库所有者负责远程发布。
