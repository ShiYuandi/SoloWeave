# 项目配置

[English](en/configuration.md) · 中文为主版本

`.soloweave/project.yaml` 是项目技术和架构选型的结构化事实来源。当前使用 `schema_version: 1`，初始化后的 `project.status` 为 `draft`（待确认）。开发者审阅配置并执行 `soloweave approve` 后，CLI 会生成 `.soloweave/decisions/ADR-NNNN.md` 和批准摘要。日后修改了已批准的关键字段，`soloweave check` 会报告不一致；有意变更时需再次审阅并执行 `approve`。

纯后端项目示例：

```yaml
schema_version: 1
project:
  name: demo-api
  type: backend
  status: draft
  team_size: 1
backend:
  language: go
  framework: custom
agents: [codex, claude, cursor]
```

`project.type` 可设为 `frontend`、`backend` 或 `fullstack`，并须填写对应技术栈。纯前端项目不能同时声明独立后端，纯后端项目同理。框架名称可自定义；`workspace` 路径必须位于项目内。JSON Schema 检查字段和值，Go 代码进一步检查跨字段组合和路径。

`init --from FILE` 可以读取已有 YAML，但即使输入写着 `approved`，初始化仍会保存为待确认状态，并要求另行批准。Agent Skill 在进行较大开发前应读取这份配置。具体依赖版本以目标项目的包清单和锁文件为准。

`soloweave catalog` 列出可选预设；`init --preset ID --name NAME` 用预设生成待确认配置。目前包括 TypeScript 前端、Go API 和 TypeScript 全栈示例。预设只是起点，自定义框架仍可使用。SoloWeave 会拒绝少数已知不兼容组合，例如 Go 后端使用 NestJS、非 JavaScript/TypeScript 前端使用 Next.js，或非 JavaScript/TypeScript 后端使用 Prisma。
