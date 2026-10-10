---
name: project-continuity
description: Use when recording a development checkpoint, handing a SoloWeave project to another AI account or coding agent, or resuming after lost chat context.
slug: shiyuandi-soloweave-project-continuity
version: 0.1.0-preview
displayName: SoloWeave 项目交接
summary: 记录真实开发进度，让新会话或不同编程 Agent 能接续项目。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# Project continuity

On resume, read `.soloweave/project.yaml`, `context/PROJECT.md`, `context/STATUS.md`, and `context/HANDOFF.md`; run `soloweave context resume` when available. Check Git and relevant source files. If they disagree with the handoff, report that the handoff is stale and use the files as evidence.

At a meaningful checkpoint, update STATUS with completed, active, blocked, and planned work; update CHANGES for important completed edits. Run `soloweave context checkpoint` with a concise task summary, next step, and only verification actually performed. Keep HANDOFF short enough for a new agent to act on without the old conversation. Do not put credentials, tokens, private keys, or real secret values in context files. If Git is unavailable, state that clearly.
