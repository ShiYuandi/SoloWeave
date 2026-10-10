# Develop with SoloWeave Skills

[简体中文](../skill-only-workflow.md)

After installing the [six new Skills](skills.md), describe development tasks to your agent normally. Routine work does not require the SoloWeave CLI or a slash command in every conversation. The agent should inspect existing rules, source, and handoff notes, then choose an appropriate workflow. Automatic selection depends on the client; explicitly name `soloweave` if it does not trigger.

## First visit to a project

The agent reads the README, dependencies, relevant code, and tests. If `.soloweave/` exists, it checks `PROJECT.md`, `STATUS.md`, `HANDOFF.md`, and ADRs against current files and Git. An existing project without SoloWeave records keeps its current conventions. Create only the minimal context files when work will continue over time or across sessions. The developer approves major architecture choices before an ADR records them; pending proposals remain pending.

## Choosing technologies for a new project

Describe what you want to build, where it will run, and any choices already made. The agent should identify the project type and consequential gaps, then discuss applicable technology and architecture dimensions. A separated Web project might consider frontend, backend, data, and delivery separately; mobile, desktop, CLI, and data projects have different concerns. Types can be combined or extended beyond the guide's examples.

The agent should offer independently combinable candidates with tradeoffs and a reasoned recommendation. You can choose each part, ask for more options, or suggest technology absent from the guide. **Examples do not limit the agent's choices**; there are no preset bundles. The agent checks compatibility, and records major decisions in an ADR only after your confirmation. It should not repeat choices you already supplied or migrate an existing project merely because the Skill was installed.

## Normal development and handoff

For a feature, search for suitable existing implementations. For a bug, seek observable reproduction evidence first. Run relevant available project checks and distinguish passed, failed, not run, and unavailable. After meaningful work, a blocker, or before switching accounts, the agent updates `STATUS.md`, `CHANGES.md`, and a short `HANDOFF.md` under `.soloweave/context/`. A new agent reads these notes and then checks Git and source files; current files win when they conflict with old notes.

`PROJECT.md` holds stable goals and approved decision links. `HANDOFF.md` should identify the current task, recent files, actual verification and outcome, open issues, next step, and check time. If Git is unavailable, say so. Do not store secrets. Skills-only work does not create the Go CLI's `approval.json`, `checkpoint.json`, or `installation.json`, or call manual review a passing `soloweave check`.

See [cross-agent acceptance](agent-acceptance.md) for the test procedure.
