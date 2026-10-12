# Public installation and full-project handoff acceptance

[简体中文](../end-to-end-acceptance.md)

Date: 2026-10-12, Asia/Shanghai. Source: pushed commit `37c7062`. No Skill rules or global installations changed in this round.

## Distribution

skills@1.7.1 discovered exactly five Skills from public GitHub. Installation into a fresh disposable project for Codex, Claude Code and Cursor produced 12 byte-matching files across two platform directories. This establishes installation artifacts, not three-client live behavior or conflict handling for user edits.

All five current SkillHub versions are published. Six source resources across five public ZIPs match the repository. Platform-generated _meta.json slug/version fields were checked separately. The retired debugging listing is unpublished and recoverable. See the [delivery snapshot](../release-status.md).

## Three fresh Codex sessions

Codex CLI 0.162.0-alpha.17.2 used configured model gpt-6.1-sol without an override. Each ephemeral session retained project files but had no prior conversation.

1. An ordinary request for a local, separated frontend/backend todo app caused actual entry/setup/feature/continuity Skill reads, dimensional recommendations and pending records. No application or approved ADR existed before approval.
2. The test driver approved native JavaScript, Node built-in HTTP and JSON storage, then explicitly scoped this phase to add/complete/persistence, leaving deletion for the next session. The agent implemented that scope, recorded an ADR and four context documents, and handed off deletion with concrete reuse instructions. Direct service and restart checks passed. The automatic integration runner failed with spawn EPERM and was accurately recorded as blocked. External execution of the preserved first-phase snapshot passed both integration tests.
3. A fresh session received only “请继续这个项目”. It read project records, ADR, code and the entry/continuity/feature Skills, then implemented the pending deletion without asking a vague direction question. It reused existing helpers, preserved the ADR and updated the records. Direct service checks and real restart recovery passed. Original automatic tests remained blocked and that session had no browser surface. External execution of the final snapshot passed both integration tests.

The external browser check actually added, completed and deleted tasks, closed the tab, stopped/restarted the backend and opened a new tab. The completed task remained, the deleted task did not return, and JSON persistence matched. Both dedicated test services were stopped.

This scenario passed within that scope. It does not establish every Skill's independent behavior, all project types, mobile layout or browser failure-path interactions. Quality review was checked for installation/distribution only in this round. Internal failures and external successes remain separate.

The [reviewed receipt](../codex-evidence/2026-10-12-public-five.json) records prompts, local Skill reads, original log hashes, snapshots and external verification. Raw evidence remains in ignored local dist/acceptance/public-five-2026-10-12/. Earlier evidence was preserved.

## Other clients

No callable claude or cursor command was found. Their live behavior remains unavailable/pending; installation artifacts and earlier Claude Code results using glm-5.3 do not establish current client or Anthropic model behavior. On a machine with those clients, repeat the three phases, record client/model/Skill versions and actual file reads, and distinguish ordinary requests from explicit invocation. See the [acceptance guide](agent-acceptance.md).
