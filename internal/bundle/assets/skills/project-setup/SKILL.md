---
name: project-setup
description: Use when creating a new SoloWeave project or adopting an existing project, selecting its technology stack, and recording approved architecture decisions.
slug: shiyuandi-soloweave-project-setup
version: 0.1.0
displayName: SoloWeave 项目规划
summary: 为新项目或现有项目确认技术选型，并记录开发者批准的架构决策。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# Project setup

For an existing project, inspect source, dependency manifests, and tests before suggesting a configuration. Do not migrate its architecture as part of adoption. For a new project, clarify the product goal and MVP, then present a small set of fitting stack and architecture choices with tradeoffs. The developer makes the final selection.

Use `soloweave init` to create a draft `.soloweave/project.yaml`. Review it with the developer before `soloweave approve`, which records the decision in an ADR. Do not silently approve or change a previously approved choice. Keep technology versions in each project's dependency files, not in the SoloWeave contract.
