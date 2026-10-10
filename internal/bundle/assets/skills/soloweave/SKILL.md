---
name: soloweave
description: Use when starting or resuming work in a SoloWeave project and the task needs routing to project setup, feature development, continuity, or quality review.
slug: shiyuandi-soloweave
version: 0.1.0-preview
displayName: SoloWeave 项目入口
summary: 在 Codex、Claude Code 和 Cursor 中延续项目决策、开发任务与交接信息。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# SoloWeave

This skill is part of a five-skill set. Install `project-setup`, `feature-workflow`, `project-continuity`, and `quality-review` alongside it. The workflows also use the separate `soloweave` CLI; get the Windows x64 program and setup instructions from https://github.com/ShiYuandi/SoloWeave/releases. If the CLI is unavailable, explain which step could not run instead of claiming it succeeded.

Read `.soloweave/project.yaml` for approved engineering decisions and `.soloweave/context/HANDOFF.md` for the current task. Check relevant source files and Git state before trusting the handoff. If no project contract exists, use `project-setup`. For a feature or fix, use `feature-workflow`; for a pause or agent switch, use `project-continuity`; for verification, use `quality-review`.

The developer chooses major architecture changes. Do not treat a draft configuration as approved. Keep the response proportional to the task; a small edit need not invoke the full workflow.
