# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is an AI engineering skills toolkit and CLI for solo developers and teams of 1–5. It keeps architecture decisions, task progress, and handoff notes in project files, so work can continue after an account change or a switch between Codex, Claude Code, and Cursor.

## Why I built SoloWeave

While using AI to write code, I found that getting one feature built is easy compared with keeping an entire project coherent over time. When I start a full-stack project from scratch, an AI agent may choose the stack and directory structure before I have weighed the options. Later, it may introduce a different approach or rewrite functionality that already exists.

Losing context is even more frustrating. When I switch accounts, models, or coding tools, the next agent does not know the project's goal, approved architecture, completed work, or open problems. I have to explain the project again, sometimes starting with another full review of the code.

I began to ask why that knowledge should live only in a chat history. If decisions, progress, code changes, actual verification results, and next steps live with the project, a new agent can check the files and continue from there.

That is why I built SoloWeave. I want AI to discuss technical choices with me before implementation, follow the decisions I approve, reuse suitable code, and leave a useful handoff after important work. The developer keeps control of major decisions, and the workflow should stay proportionate to the project.

SoloWeave started with my own development needs. I am sharing it so other independent developers and small teams can keep projects moving across AI tools, instead of starting over with each conversation.

## Core capabilities

- **Developer-approved architecture**: record technical choices and approve key decisions before implementation.
- **Project continuity**: keep progress, changes, actual verification results, and next steps with the project.
- **Consistent engineering**: help agents inspect existing code and follow approved project conventions.

The **v0.1.0-preview Windows x64 build** includes the CLI, a YAML project contract that requires developer approval, five Agent Skills, platform rule installation, and a handoff workflow. It does not call a remote AI API or require a database. The [v0.2 design baseline](../design-v0.2.md) and [development plan](../development-plan.md) are maintained in Chinese.

## Download and run on Windows

1. [Download the v0.1.0-preview Windows x64 package](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/soloweave-v0.1.0-preview-windows-amd64.zip). See [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) for other versions.
2. Extract the ZIP to get `soloweave.exe`. **You do not need Go or a local build.**
3. Run the quick-start commands below in PowerShell.

Choose `soloweave-...-windows-amd64.zip` under **Assets**. GitHub's automatic `Source code (zip)` and `Source code (tar.gz)` archives contain source files, not a ready-to-run executable. The package also includes Chinese and English usage guides and the MIT license. Use [SHA256SUMS](https://github.com/ShiYuandi/SoloWeave/releases/download/v0.1.0-preview/SHA256SUMS) to check the downloaded archive. Only Windows x64 is packaged today.

## Quick start

Add the extracted directory to `PATH`, then run these commands inside the target project. Replace the example paths with your own in PowerShell:

```powershell
$env:PATH = "C:\Tools\SoloWeave;$env:PATH"
cd C:\path\to\your-project
```

Then run:

```sh
soloweave init
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave check
```

`init` offers stack presets and a custom setup. Run `soloweave catalog` to list presets. For non-interactive use, try `init --preset api-go --name demo` or `init --from path/to/project.yaml`. Both create a **draft** and still require a separate `approve`. Use `approve --yes` only after reviewing the configuration.

Before pausing a task or switching agents, record the actual progress and verification:

```sh
soloweave context checkpoint --task "login" --summary "form complete" --next "connect API" --verification "go test ./...: passed"
soloweave context resume
soloweave context check
```

Omit `--verification` if no check ran; the handoff will say `Not run`. Without a Git repository, Git state is reported as `UNAVAILABLE`. See [project continuity](project-continuity.md) and [configuration](configuration.md).

## Installed files

Project configuration, ADRs, and handoff files live in `.soloweave/`. Codex and Cursor share `.agents/skills/`; Claude Code uses `.claude/skills/`. Platform rules go to `AGENTS.md`, `CLAUDE.md`, and `.cursor/rules/soloweave.mdc`. The installer previews changes and stops on conflicting user files instead of silently overwriting them.

See the [repository directory guide](directory-structure.md) and [Skills guide](skills.md).

## Verification status

Go tests, vet, build, validation of the five Skills, and a complete Windows CLI smoke test in a temporary project have passed. Linux and macOS binaries have only been cross-compiled. Live handoffs in Claude Code and Cursor remain untested; see the [acceptance procedure](agent-acceptance.md). For source builds and development, see the [contribution guide](CONTRIBUTING.md).

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
