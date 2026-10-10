# Skills-first acceptance

[简体中文](../agent-acceptance.md)

This page separates **local file and packaging checks** from **live agent behavior**. Skill descriptions help agents select them but do not replace automatic invocation tests.

## Local checks

1. Validate names, frontmatter, versions, and unique slugs in the six root `SKILL.md` files. Run `scripts/package-skillhub.ps1` and inspect ZIP contents and SHA256SUMS.
2. Provide the six Skills to a disposable project without the SoloWeave Go CLI. Keep a reusable function and related test, then give the agent an ordinary development request. Check that it reads the project, reuses code, and runs a real test.
3. Propose a major technical choice without approving it. Check that records stay pending and no ADR or CLI approval file is fabricated. After explicit approval, check the ADR records the decision and reasoning.
4. Resume in another session without the old chat. Confirm the agent finds the goal, current task, recent changes, verification, and next step from project files. After changing a relevant source file, check it recognizes that old handoff notes may be stale.
5. Regress a blocked new project: have the agent implement a project that needs local build scripts, then cause an observable dependency or permission blocker. Before window 1 ends, verify that `STATUS.md` and `HANDOFF.md` record the actual failure, blocker, and next step. In window 2, ask only “continue this project”; check that the agent verifies and advances or explains the blocker, without calling unverified work runnable. Any extra authorization must be requested for a concrete action.

## Live client checks

Install the new Skills in disposable Codex, Claude Code, and Cursor projects. Record each client version, installation path, and result. Test ordinary-request automatic invocation, developer confirmation of major decisions, automatic handoff after a feature or fix, and verification of actual code in a new session. If automatic invocation fails, record that separately from success after explicitly naming the entry Skill.

Record public GitHub `--list` discovery, real installation, and SkillHub publication separately. Local folder discovery is not proof of public installation. If needed, test the legacy CLI v0.1.1-preview using its historical workflow.

### Project selection regression (`project-setup` 0.3.0)

Use these prompts in a fresh session with the revised Skill installed. Record its path and version, the original prompt, the agent response, project files, and observable invocation status. Static file checks do not establish live behavior.

| Prompt or condition | Expected behavior |
| --- | --- |
| “Build a simple persistent todo app in an empty directory.” | Clarify a project-shape decision such as whether a separate backend is needed before assuming a static page. |
| “Build a separated frontend/backend todo app; recommend technologies.” | Compare applicable frontend, backend, data, architecture, and delivery dimensions with independently combinable choices and reasons, not three bundled stacks. |
| “Build a mobile app with a backend API.” | Cover mobile and API concerns, check their interface and delivery, and avoid assuming one or multiple repositories. |
| “Build a cross-platform file-organizing CLI.” | Cover runtime, distribution, configuration, and failure handling without forcing Web frameworks. |
| “Analyze offline sales data without ML models.” | Cover data and processing without forcing model/provider choices. |
| “I want technology X, which is absent from your guide.” | Evaluate X against requirements rather than rejecting it as unlisted; check compatibility where needed. |
| Continue an existing project with dependencies and an ADR. | Verify actual implementation and effective decisions; do not migrate because an example differs. |
| Provide every important technology and architecture choice up front. | Check the combination and restate what needs confirmation without asking the same questions again. |
| Propose a combination incompatible with the target platform. | Explain the specific conflict and alternatives; do not implement it or write an approved ADR before confirmation. |
| Ask about a version or compatibility that cannot be verified. | State uncertainty and what to verify; do not invent current support. |

Also try a desktop project and one unlisted project type to check that the agent generates relevant dimensions. If the client does not expose invocation events, record invocation as unobservable rather than inferring it from the reply.

## Current results

The five earlier Skills were exercised in a CLI-free disposable project. For the six new Skills, the repository packaging script checked names, required fields, versions, and unique slugs; six ZIP entries and SHA256SUMS were verified. The standard installer's `add . --list` discovered all six Skills. In disposable projects, its Codex, Claude Code, and Cursor project-level installations completed, and each installed `SKILL.md` matched the source. A default-mode Codex installation on Windows also completed.

In a separate project with neither the SoloWeave CLI nor Git, a manual exercise followed the new Skill workflow: a test first failed on a missing module, then passed two checks after a small implementation reused an existing function. Project, status, change, and handoff records were written. A separate read and rerun recovered the task context and reproduced the passing result. This exercise establishes that the file workflow is usable; it does not establish automatic Skill invocation by an agent.

The legacy Go CLI's `go test ./... -count=1`, `go vet ./...`, and `go build ./...` exited successfully. Relative file links in changed documentation passed a local check.

The `quick_validate.py` attempt did not complete because PyYAML is missing, so that validator is not reported as passing. An initial standard-installer attempt hit an npm connection refusal; a later retry completed the local discovery and installation checks above.

### User-provided two-window Codex trial (2026-10-10)

- Window 1 reported installing all six Skills into Codex's global directory from public GitHub commit `5bc8f1e`. This review confirmed that all six installed `SKILL.md` files have the same SHA256 hashes as the repository sources. An ordinary todo-app request led to an architecture proposal; after the developer requested a separated frontend and backend and approved the choice, the agent wrote source files, `PROJECT.md`, and an ADR. The conversation excerpt alone does not prove that the client automatically invoked a Skill internally.
- Window 1 reported that local pnpm policy blocked the `better-sqlite3` and `esbuild` build scripts, so a full build had not completed. Inspection of the disposable `test/` project outside this repository found only `PROJECT.md` and an ADR, with no `STATUS.md`, `HANDOFF.md`, or `CHANGES.md`; its README did not record that startup blocker.
- In window 2, the ordinary request “continue this project” received a claim that the app was runnable and a question about which direction to take. It did not address the known startup blocker or use a handoff checkpoint. The cross-session continuity check **failed**. This reply alone cannot establish whether window 2 loaded or invoked the Skill; automatic invocation still needs observable client evidence.

The repository sources of `soloweave`, `feature-workflow`, and `project-continuity` have since been revised to `0.2.1` to require a checkpoint when blocked and resumption from known unfinished work. **The revision has not yet been retested in a live client.** The user's global installation still contains the `0.2.0` trial version and must be updated before the two-window regression.

Separately, `project-setup` is now `0.3.0` with an on-demand selection guide. The local packaging run succeeded; checks of the six current ZIP files verified SHA256SUMS, entries, and byte-for-byte source matches, including both `SKILL.md` and the guide in the `project-setup` ZIP. The standard installer `skills@1.7.1 add . --list` discovered all six locally. A copied installation in an isolated project contained both `project-setup` files with matching SHA256 hashes. **The new selection behavior has not been retested in a fresh live agent session.** Static and installation checks do not establish agent behavior. `quick_validate.py` still cannot run because PyYAML is unavailable here.

**Public-GitHub standard-installer commands, SkillHub updates, and live Claude Code/Cursor behavior remain independently unverified.**
