---
name: feature-workflow
description: Use when implementing a feature or bug fix in a SoloWeave project while preserving its approved architecture and reusing existing code.
slug: shiyuandi-soloweave-feature-workflow
version: 0.1.0
displayName: SoloWeave 功能开发
summary: 在已确认的架构下检查现有代码、复用合适实现并完成需求或修复。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# Feature workflow

Read the approved project contract and current handoff. Inspect related modules, existing components, utilities, dependencies, and tests before designing a change. Reuse a suitable implementation; create new code when reuse would increase coupling or obscure responsibility. Keep the change as small as the requirement permits.

State a short implementation plan for substantial work. Make the edit, run relevant verification, and report what actually ran. If the task changes an approved technology or architecture decision, pause that change for developer approval and a new ADR. At meaningful checkpoints update the project status and handoff using `project-continuity`.
