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

## Five Skills: roles and when to use them

Each Skill covers one part of development. **Install the full set or select individual Skills.** The agent should choose what fits the task; it does not need to run all five every time.

| Skill | What it does | When it applies |
| --- | --- | --- |
| **Entry** [soloweave](../../skills/soloweave/SKILL.md) | Reads the directory, rules, and project records; distinguishes new and existing projects and selects a workflow. | Starting development, planning, code changes, review, or asking to continue a project. |
| **Setup** [project-setup](../../skills/project-setup/SKILL.md) | Clarifies requirements and compares relevant technology and architecture choices; records major decisions and reasons after your confirmation. | Creating a project in an empty directory, establishing conventions for an existing project, or making major technical choices. Existing implementations remain the starting point. |
| **Feature development** [feature-workflow](../../skills/feature-workflow/SKILL.md) | Searches for reusable code, follows agreed conventions, implements changes, runs relevant checks, and records progress. | Adding features or changing existing code. New projects first clarify requirements and confirm major choices. |
| **Quality review** [quality-review](../../skills/quality-review/SKILL.md) | Checks actual changes, project decisions, and verification evidence; reports specific issues, locations, and remaining gaps. | Requesting code review, checking delivery quality, or assessing readiness for handoff. |
| **Continuity** [project-continuity](../../skills/project-continuity/SKILL.md) | Saves actual progress, changes, verification, and next steps; checks files before resuming known work. | Starting a new session, finishing important work, encountering a blocker, or switching accounts, models, or tools. |

### How they work together

- **“Build a todo app from scratch.”** Entry identifies a new project → setup clarifies requirements and compares choices → feature development starts after you confirm major decisions → meaningful checkpoints leave handoff records.
- **“Check whether these changes are ready to deliver.”** Quality review checks requirements, actual files, and verification results; it reports specific issues and remaining gaps, with a checkpoint when handoff is needed.
- **“I switched accounts; continue this project.”** Continuity reads goals, status, and next steps, verifies actual code and checks, then selects the applicable workflow. Remaining blockers stay explicit.

Automatic selection depends on the client and is not guaranteed. If needed, name a Skill explicitly, such as “use project-setup to help me compare technologies.” See [Skill details](skills.md).

## Quick start

The repository's [`skills/`](../../skills/) directory contains five independent Skills. Ask a compatible AI tool to install them from this repository, or use the [Skills CLI](https://github.com/vercel-labs/skills):

```sh
npx skills@latest add ShiYuandi/SoloWeave
```

Or send this prompt to your coding agent and let it handle installation:

```text
Follow https://github.com/vercel-labs/skills to install all five SoloWeave Skills from the skills/ directory of https://github.com/ShiYuandi/SoloWeave for this coding agent. Check for existing files with the same names and do not overwrite my changes. Report the path and version of each installed SKILL.md. For future ordinary development requests, use the applicable SoloWeave workflow.
```

Then ask for work normally, for example:

> Add user login to this project. Inspect existing code first, run the relevant tests, and leave a progress note another agent can use.

After installation, describe your development task normally. The entry Skill routes to setup, feature development, review, or continuity as needed. The agent should inspect an existing project's code and records; you still approve major technical choices. Automatic skill selection depends on the agent and is not guaranteed. If needed, explicitly ask it to “use SoloWeave to continue this project.”

**Verification scope:** The current five-Skill set passed 15 maintenance checks, independent installations, three ordinary-request Codex scenarios, and one explicitly invoked blocked checkpoint. See [results and limits](../five-skills-acceptance.md) and [acceptance](agent-acceptance.md). The earlier [12 Codex sessions](../codex-behavior-results.md) tested the six-Skill set and remain historical evidence. Earlier Claude Code evidence used `glm-5.3`; Cursor and Anthropic models remain untested. See [delivery status](../release-status.md) for GitHub and SkillHub synchronization.

## Where project records live

Use existing project documentation where possible. New records live in `.soloweave/context/` and `.soloweave/decisions/`. **Install Skills directly; there is no program to download or build.** Skills maintain Markdown records and use the target project's existing checks. See the [workflow](skill-only-workflow.md) and [Skill roles](skills.md).

## Verification and contribution

The [acceptance guide](agent-acceptance.md) records local validation and separates it from live client tests. See the [directory guide](directory-structure.md) and [contribution guide](CONTRIBUTING.md) for the repository and source development.

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
