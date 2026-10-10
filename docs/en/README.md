# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is an AI engineering skills toolkit and CLI for solo developers and teams of 1–5. It keeps architecture decisions, task progress, and handoff notes in project files, so work can continue after an account change or a switch between Codex, Claude Code, and Cursor.

## Why I built SoloWeave

While using AI to write code, I found that **building one feature is easier than keeping an entire project coherent over time**. When I start a full-stack project from scratch, an AI agent may **choose the stack and directory structure before I approve them**. Later, it may introduce a different approach or **rewrite functionality that already exists**.

**Losing development context** is even more frustrating. When I switch accounts, models, or coding tools, the next agent does not know the project's goal, approved architecture, completed work, or open problems. I have to **explain the project again and ask the AI to reexamine existing work**.

I began to ask why that knowledge should live only in a chat history. If **decisions, progress, code changes, actual verification results, and next steps live with the project**, a new agent can check the files and continue from there.

That is why I built SoloWeave. I want AI to **discuss technical and architectural choices with me before implementation**, **follow approved decisions and reuse suitable code**, and **leave a useful handoff after important work**. The developer keeps control of major decisions, and the workflow should stay proportionate to the project.

SoloWeave started with my own development needs. I am sharing it so other independent developers and small teams can **keep projects moving across AI tools instead of starting over with each conversation**.

## Core capabilities

- **Developer-approved architecture**: record technical choices and approve key decisions before implementation.
- **Project continuity**: keep progress, changes, actual verification results, and next steps with the project.
- **Consistent engineering**: help agents inspect existing code and follow approved project conventions.

The **v0.1.1-preview Windows x64 build** includes the CLI, a YAML project contract that requires developer approval, five Agent Skills, platform rule installation, and a handoff workflow. It does not call a remote AI API or require a database. The [v0.2 design baseline](../design-v0.2.md) and [development plan](../development-plan.md) are maintained in Chinese.

## Skills only, or use the CLI?

| Option | What you get | Current limitation |
| --- | --- | --- |
| **Install the five Skills only** | Your AI can discuss architecture, inspect and reuse code, and manually record decisions you approve, progress, and handoffs in project files. **You do not need the CLI** if you only want these conventions. | This path does not create CLI approval or installation records or automatically checkable checkpoints. CLI commands such as `soloweave check` are unavailable; the AI must report them as not run. |
| **Use the CLI (which installs the Skills)** | The same guidance, plus commands that create project and handoff files, record approved decisions, install the five Skills and project rules, and check file state. | A ready-to-run CLI package is currently available only for Windows x64. |

The CLI works with local files. It does not write application code for the AI, approve technical choices for the developer, or assume tests passed.

## Let your AI install the Skills

If your coding agent can access the internet and install Skills, send it this prompt:

```text
Follow https://skillhub.cn/install/skillhub.md to install these five SoloWeave Skills for the current coding agent:
@user_38c0807e/shiyuandi-soloweave
@user_38c0807e/shiyuandi-soloweave-project-setup
@user_38c0807e/shiyuandi-soloweave-feature-workflow
@user_38c0807e/shiyuandi-soloweave-project-continuity
@user_38c0807e/shiyuandi-soloweave-quality-review
Check for existing files with the same names first. Do not overwrite my changes. After installation, verify all five SKILL.md files and report their paths and versions.
```

The entry Skill needs the other four. This prompt installs Skill files only; **you can stop here if you only want the development conventions**. After installation, you can tell your AI: “Use SoloWeave's Skills-only workflow. Inspect the existing code first, manually record decisions I approve and handoffs, and tell me which CLI checks were not run.” See the [Skills-only guide](skill-only-workflow.md) for file locations and acceptance steps.

If you choose the CLI path below, **you do not need to install from SkillHub first**: `soloweave install` installs all five Skills and the platform rules. The CLI's `check` and `doctor` commands do not count a SkillHub installation as a CLI-managed project installation. If both methods have written to the same directory, preview and resolve any conflicts first. See the [Skills guide](skills.md).

## Download and run on Windows

1. [Download the v0.1.1-preview Windows x64 package](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.1-preview/soloweave-v0.1.1-preview-windows-amd64.zip). See [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) for other versions.
2. Extract the ZIP to get `soloweave.exe`. **You do not need Go or a local build.**
3. Run the quick-start commands below in PowerShell.

Choose `soloweave-...-windows-amd64.zip` under **Assets**. GitHub's automatic `Source code (zip)` and `Source code (tar.gz)` archives contain source files, not a ready-to-run executable. The package also includes Chinese and English usage guides and the MIT license. Use [SHA256SUMS](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.1-preview/SHA256SUMS) to check the downloaded archive. Only Windows x64 is packaged today.

## What the CLI commands do

Run these commands in the target project directory; they operate on that project's files.

| Command | Purpose |
| --- | --- |
| `soloweave catalog` | List optional stack presets. |
| `soloweave init` | Create a **draft** `.soloweave/project.yaml` and initial project and handoff documents. It does not approve the technical choices for you. |
| `soloweave approve` | Show the configuration and, after developer confirmation, mark it approved and create an architecture decision record (ADR). |
| `soloweave install --dry-run` | Preview the five Skills and platform rules without writing files. A regular `install` writes them while protecting existing or edited files. |
| `soloweave context checkpoint` | Write a handoff from a task summary, next step, and **actual** verification supplied by you or the AI. It captures Git state when available; it does not infer progress or test results. |
| `soloweave context resume` / `context show` | Read project and handoff information so a new session can take over. |
| `soloweave context check` | Check for missing or stale handoff information. |
| `soloweave check` | Check the project configuration, architecture approval, CLI installation record, and handoff state. |
| `soloweave doctor` | Report Git, project configuration, and CLI installation status to help diagnose setup issues. |
| `soloweave version` | Print the CLI version. |

## Quick start

Add the extracted directory to `PATH`, then run these commands inside the target project. Replace the example paths with your own in PowerShell:

```powershell
$env:PATH = "C:\Tools\SoloWeave;$env:PATH"
cd C:\path\to\your-project
```

Then follow this **CLI path (including the Skills)**. `--agents codex,claude,cursor` installs project files for all three platforms; list only the ones you use if you need fewer.

```sh
soloweave init
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave context checkpoint --task "project setup" --summary "configuration approved and Skills installed" --next "start development"
soloweave check
```

`init` offers stack presets and a custom setup. Run `soloweave catalog` to list presets. For non-interactive use, try `init --preset api-go --name demo` or `init --from path/to/project.yaml`. Both create a **draft** and still require a separate `approve`. Use `approve --yes` only after reviewing the configuration.

Before pausing a task or switching agents, record the actual progress and verification:

```sh
soloweave context checkpoint --task "login" --summary "form complete" --next "connect API"
soloweave context resume
soloweave context check
```

Add `--verification "command: actual result"` to `context checkpoint` only after running that check.

Omit `--verification` if no check ran; the handoff will say `Not run`. Without a Git repository, Git state is reported as `UNAVAILABLE`. See [project continuity](project-continuity.md) and [configuration](configuration.md).

## Installed files

Project configuration, ADRs, and handoff files live in `.soloweave/`. Codex and Cursor share `.agents/skills/`; Claude Code uses `.claude/skills/`. Platform rules go to `AGENTS.md`, `CLAUDE.md`, and `.cursor/rules/soloweave.mdc`. The installer previews changes and stops on conflicting user files instead of silently overwriting them.

See the [repository directory guide](directory-structure.md) and [Skills guide](skills.md).

## Verification status

Go tests, vet, build, validation of the five Skills, and a complete Windows CLI smoke test in a temporary project have passed. Linux and macOS binaries have only been cross-compiled. Live handoffs in Claude Code and Cursor remain untested; see the [acceptance procedure](agent-acceptance.md). For source builds and development, see the [contribution guide](CONTRIBUTING.md).

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
