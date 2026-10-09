# Project configuration

[简体中文](../configuration.md)

`.soloweave/project.yaml` is the structured source of truth for a project's technology and architecture choices. The current format uses `schema_version: 1`; initialization sets `project.status: draft`. After the developer reviews the file and runs `soloweave approve`, the CLI writes `.soloweave/decisions/ADR-NNNN.md` and an approval digest. If a critical approved field later changes, `soloweave check` reports the mismatch. Review and approve an intentional change again.

Example backend project:

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

`project.type` can be `frontend`, `backend`, or `fullstack`; the corresponding stack sections are required. A frontend-only project cannot declare an independent backend, and vice versa. Framework names are open strings, so custom stacks are valid. `workspace` paths must remain within the project. JSON Schema checks fields and values; Go code checks cross-field combinations and paths.

`init --from FILE` reads an existing YAML file, but even an input marked `approved` is saved as a draft and requires separate approval. Agent Skills should read the contract before substantial development. Dependency versions remain in the target project's manifests and lock files.

`soloweave catalog` lists optional presets; `init --preset ID --name NAME` creates a draft from one. Current examples include a TypeScript frontend, Go API, and TypeScript fullstack app. Presets are starting points, not a restriction on custom frameworks. SoloWeave rejects a few known incompatible combinations, such as NestJS with a Go backend, Next.js with a non-JavaScript frontend, or Prisma without a JavaScript/TypeScript backend.
