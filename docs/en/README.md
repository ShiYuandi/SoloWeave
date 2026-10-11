# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is a set of AI development Skills for solo developers and teams of 1–5. After installation, ask for development work normally. The agent should read project facts, follow approved decisions, look for reusable code, and maintain a useful handoff after meaningful work.

## Why I built SoloWeave

While using AI to write code, I found that **building one feature is easier than keeping an entire project coherent over time**.

When starting a project, AI usually proposes technologies and asks me to confirm them. But I may still not understand **why those technologies were recommended, which alternatives fit, or what the tradeoffs are**. I want it to **understand my requirements first, then discuss relevant choices such as frontend, backend, database, and architecture separately**, so I can make an informed decision.

As development continues, AI may also drift from agreed conventions or overlook existing components and functions, **rewriting functionality that already exists**.

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

After installation, describe your development task normally. The entry Skill routes to setup, feature development, debugging, review, or continuity as needed. The agent should inspect an existing project's code and records; you still approve major technical choices. Automatic skill selection depends on the agent and is not guaranteed. If needed, explicitly ask it to “use SoloWeave to continue this project.”

**Verification status:** Independent installation file checks, static checks, and packaging passed. Current source versions also passed six standalone minimal scenarios, empty-directory routing, and blocked handoff in real Codex CLI sessions. See the [12-session results](../codex-behavior-results.md) and [acceptance](agent-acceptance.md). Earlier Claude Code evidence used older revisions and the actual model `glm-5.3`; Cursor and Anthropic models were not tested in this round. Actual GitHub and SkillHub progress is recorded in [delivery status](../release-status.md).

## What it does

| When | Expected agent behavior |
| --- | --- |
| Enter an existing project | Read rules, goals, decisions, status, and code; check whether the handoff is stale. |
| Create or adopt a project | Route an empty directory through project setup first. Understand requirements, compare independently combinable choices, and obtain your approval before implementing major decisions or recording an ADR. Examples do not limit the agent's choices. |
| Develop or fix code | Search for suitable existing implementations, make the change, and run relevant checks. |
| Finish important work or switch accounts | Update progress, significant changes, actual test results, and next steps for the next session. |

Use existing project documentation where possible. New records live in `.soloweave/context/` and `.soloweave/decisions/`. **Install Skills directly; there is no program to download or build.** Skills maintain Markdown records and use the target project's existing checks. See the [workflow](skill-only-workflow.md) and [Skill roles](skills.md).

## Verification and contribution

The [acceptance guide](agent-acceptance.md) records local validation and separates it from live client tests. See the [directory guide](directory-structure.md) and [contribution guide](CONTRIBUTING.md) for the repository and source development.

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
