# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is an AI engineering skills toolkit and CLI for solo developers and teams of 1–5. It keeps architecture decisions, task progress, and handoff notes in project files, so work can continue after an account change or a switch between Codex, Claude Code, and Cursor.

**V0.1 is a local preview.** It includes a Go CLI, a YAML project contract that requires developer approval, five Agent Skills, platform rule installation, and a handoff workflow. It does not call a remote AI API or require a database. The [v0.2 design baseline](../design-v0.2.md) and [development plan](../development-plan.md) are maintained in Chinese.

## Build and test

Use Go 1.27.2 or newer from the repository root:

```sh
go test ./...
go vet ./...
go build -o soloweave ./cmd/soloweave
```

On Windows, use `soloweave.exe` as the output name. `scripts/build-preview.ps1` creates local Windows, Linux, and macOS preview binaries plus SHA256 checksums in the Git-ignored `dist/` directory. Cross-compilation does not verify runtime behavior on Linux or macOS.

## Quick start

Put the compiled `soloweave` executable on your `PATH`, then run these commands inside the target project:

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

Go tests, vet, build, validation of the five Skills, and a complete Windows CLI smoke test in a temporary project have passed. Linux and macOS binaries have only been cross-compiled. Live handoffs in Claude Code and Cursor remain untested; see the [acceptance procedure](agent-acceptance.md). Remote CI and a public Release have not run.

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
