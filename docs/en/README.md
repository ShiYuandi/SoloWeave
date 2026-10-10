# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is a set of AI development Skills for solo developers and teams of 1–5. After installation, ask for development work normally. The agent should read project facts, follow approved decisions, look for reusable code, and maintain a useful handoff after meaningful work.

## Why I built SoloWeave

While using AI to write code, I found that **building one feature is easier than keeping an entire project coherent over time**. An agent may **choose a stack or directory structure before I approve it**, then introduce another approach later or **rewrite functionality that already exists**.

**Losing development context** is even more frustrating. After I switch accounts, models, or tools, the next agent does not know the project's goal, approved architecture, completed work, or open problems. I have to **explain the project again and ask the AI to reexamine existing work**.

I began to ask why that knowledge should live only in chat history. If **decisions, progress, code changes, actual verification results, and next steps live with the project**, another agent can check the files and continue from there.

That is why I built SoloWeave. I want AI to **discuss important technical choices with me before implementation**, **follow approved decisions and reuse suitable code**, and **leave a clear handoff after important work**. The developer retains control of major decisions, and the process should stay proportionate to the task.

I am sharing it so other independent developers and small teams can **keep projects moving across AI tools instead of starting over with each conversation**.

## Quick start

The repository's [`skills/`](../../skills/) directory contains six independent Skills. Ask a compatible AI tool to install them from this repository, or use the [Skills CLI](https://github.com/vercel-labs/skills):

```sh
npx skills@latest add ShiYuandi/SoloWeave
```

Or send this prompt to your coding agent and let it handle installation:

```text
Follow https://github.com/vercel-labs/skills to install all six SoloWeave Skills from the skills/ directory of https://github.com/ShiYuandi/SoloWeave for this coding agent. Check for existing files with the same names and do not overwrite my changes. Report the path and version of each installed SKILL.md. For future ordinary development requests, use the applicable SoloWeave workflow.
```

Then ask for work normally, for example:

> Add user login to this project. Inspect existing code first, run the relevant tests, and leave a progress note another agent can use.

You do not need to type a SoloWeave command for every task. The entry Skill routes to setup, feature development, debugging, review, or continuity as needed. The agent should inspect an existing project's code and records; you still approve major technical choices. Automatic skill selection depends on the agent and is not guaranteed. If needed, explicitly ask it to “use SoloWeave to continue this project.”

**Verification status:** The standard installer discovered all six new Skills from the local repository and created the corresponding project Skill files for Codex, Claude Code, and Cursor in a disposable project. Live automatic invocation, installation from the public GitHub URL, and SkillHub updates still need separate acceptance. The current [SkillHub pages](skills.md#skillhub-publication-status) host the five previously published Skills; check their displayed versions.

## What it does

| When | Expected agent behavior |
| --- | --- |
| Enter an existing project | Read rules, goals, decisions, status, and code; check whether the handoff is stale. |
| Create or adopt a project | Inspect what exists, explain key options, obtain your approval for major choices, and record an ADR. |
| Develop or fix code | Search for suitable existing implementations, make the change, and run relevant checks. |
| Finish important work or switch accounts | Update progress, significant changes, actual test results, and next steps for the next session. |

Use existing project documentation where possible. New records live in `.soloweave/context/` and `.soloweave/decisions/`. **Go CLI is not required** for the Skills path. Skills do not create the CLI's machine approval, installation, or checkpoint records and do not claim to have run CLI checks. See the [Skills-only guide](skill-only-workflow.md) and [Skill roles](skills.md).

## Optional legacy CLI

The previously released [v0.1.1-preview Windows x64 package](https://github.com/ShiYuandi/SoloWeave/releases/tag/v0.1.1-preview) remains available without a local build. It provides `init` (draft configuration), `approve` (developer approval and ADR), `install` (the five Skills embedded at release time), `context checkpoint/resume/check` (handoffs), `check` (project validation), and `doctor` (environment diagnostics). The executable is a release snapshot and **does not automatically include the new workflows under `skills/`**. Existing users can keep using it; new projects should start with the Skills path above. See the [legacy configuration guide](configuration.md) and [continuity guide](project-continuity.md) for command details.

| Command | Purpose |
| --- | --- |
| `catalog` | List legacy technology presets. |
| `init` / `approve` | Create a draft; after review, record developer approval and an ADR. |
| `install --dry-run` / `install` | Preview or install embedded legacy Skills and project rules. |
| `context checkpoint` / `resume` / `check` | Save a task summary and actual verification, read a handoff, or check staleness. |
| `check` / `doctor` / `version` | Validate project files, diagnose the environment, or show the executable version. |

## Verification and contribution

The [acceptance guide](agent-acceptance.md) records local validation and separates it from live client tests. See the [directory guide](directory-structure.md) and [contribution guide](CONTRIBUTING.md) for the repository and source development.

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
