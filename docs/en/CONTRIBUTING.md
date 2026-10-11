# Contributing

[简体中文](../../CONTRIBUTING.md)

Read the [Skills design](../design.md), [development plan](../development-plan.md), and [handoff](../../.soloweave/context/HANDOFF.md), then verify actual files and Git state.

Users install Skills directly. Maintainers need **PowerShell 7** for repository scripts; No application build is required.

- Edit Skills only in root `skills/`. Keep instructions focused and references within each Skill.
- `name`, `description`, and `license` are Agent Skills fields. Preserve `slug`, `version`, `displayName`, `summary`, and `homepage` as SkillHub extensions. Bump the numeric version of each changed Skill independently.
- Verify meaningful behavior changes in real clients. Static validation and installation do not prove invocation.
- Update project status and handoff after important changes. Record developer-approved architecture in an ADR.
- Commit messages must be Chinese. Commits, pushes, and publication require an explicit developer instruction.

Run from the repository root:

```powershell
pwsh -NoProfile -File scripts/check-skills.ps1
pwsh -NoProfile -File scripts/check-docs.ps1
pwsh -NoProfile -File scripts/test-tooling.ps1
pwsh -NoProfile -File scripts/package-skillhub.ps1
```

The validator checks this repository's single-line metadata format and resource boundaries; it is not a general YAML parser. Tests cover invalid metadata, missing or escaping references, ZIP contents, SHA256, and reproducible packaging. Output goes to ignored `dist/skills/`.

See the [release guide](releasing.md) and [acceptance guide](agent-acceptance.md). Do not commit caches, artifacts, credentials, or unreviewed local session logs.

Use [live behavior testing](behavior-testing.md) for isolated fixtures and an authenticated Codex CLI. This is separate from static CI, which does not require Node.js.
