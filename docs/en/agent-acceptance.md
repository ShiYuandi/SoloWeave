# Skills-first acceptance

[简体中文](../agent-acceptance.md)

This page separates **local file and packaging checks** from **live agent behavior**. Skill descriptions help agents select them but do not replace automatic invocation tests.

## Local checks

1. Validate names, frontmatter, versions, and unique slugs in the six root `SKILL.md` files. Run `scripts/package-skillhub.ps1` and inspect ZIP contents and SHA256SUMS.
2. Provide the six Skills to a disposable project without the SoloWeave Go CLI. Keep a reusable function and related test, then give the agent an ordinary development request. Check that it reads the project, reuses code, and runs a real test.
3. Propose a major technical choice without approving it. Check that records stay pending and no ADR or CLI approval file is fabricated. After explicit approval, check the ADR records the decision and reasoning.
4. Resume in another session without the old chat. Confirm the agent finds the goal, current task, recent changes, verification, and next step from project files. After changing a relevant source file, check it recognizes that old handoff notes may be stale.

## Live client checks

Install the new Skills in disposable Codex, Claude Code, and Cursor projects. Record each client version, installation path, and result. Test ordinary-request automatic invocation, developer confirmation of major decisions, automatic handoff after a feature or fix, and verification of actual code in a new session. If automatic invocation fails, record that separately from success after explicitly naming the entry Skill.

Record public GitHub `--list` discovery, real installation, and SkillHub publication separately. Local folder discovery is not proof of public installation. If needed, test the legacy CLI v0.1.1-preview using its historical workflow.

## Current results

The five earlier Skills were exercised in a CLI-free disposable project. For the six new Skills, the repository packaging script checked names, required fields, versions, and unique slugs; six ZIP entries and SHA256SUMS were verified. The standard installer's `add . --list` discovered all six Skills. In disposable projects, its Codex, Claude Code, and Cursor project-level installations completed, and each installed `SKILL.md` matched the source. A default-mode Codex installation on Windows also completed.

In a separate project with neither the SoloWeave CLI nor Git, a manual exercise followed the new Skill workflow: a test first failed on a missing module, then passed two checks after a small implementation reused an existing function. Project, status, change, and handoff records were written. A separate read and rerun recovered the task context and reproduced the passing result. This exercise establishes that the file workflow is usable; it does not establish automatic Skill invocation by an agent.

The legacy Go CLI's `go test ./... -count=1`, `go vet ./...`, and `go build ./...` exited successfully. Relative file links in changed documentation passed a local check.

The `quick_validate.py` attempt did not complete because PyYAML is missing, so that validator is not reported as passing. An initial standard-installer attempt hit an npm connection refusal; a later retry completed the local discovery and installation checks above. **Live automatic invocation, public GitHub installation, SkillHub updates, and cross-client handoffs for the new six Skills remain unverified.**
