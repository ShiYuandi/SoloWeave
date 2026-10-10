# SoloWeave

[简体中文](../../README.md)

**Build independently. Ship confidently.**

SoloWeave is an AI engineering skills toolkit and CLI for solo developers and teams of 1–5. It keeps architecture decisions, task progress, and handoff notes in project files, so work can continue after an account change or a switch between Codex, Claude Code, and Cursor.

**V0.1 is a local preview.** It includes a Go CLI, a YAML project contract that requires developer approval, five Agent Skills, platform rule installation, and a handoff workflow. It does not call a remote AI API or require a database. The [v0.2 design baseline](../design-v0.2.md) and [development plan](../development-plan.md) are maintained in Chinese.

## Download for Windows

Download `soloweave-v*-windows-amd64.zip` from [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) under **Assets**. Extract it and run `soloweave.exe`; **Go and a local build are not required**. If no package is listed, the first release is still being prepared. The ZIP also includes Chinese and English READMEs and the MIT license. Use `SHA256SUMS` to check the downloaded archive.

The current package targets Windows x64. Maintainers can create and verify the same ZIP locally using the [contribution guide](CONTRIBUTING.md).

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

## Build and test for contributors

Source development requires Go 1.27.2 or newer. Users of the downloaded executable do not need Go. The module path is `github.com/ShiYuandi/SoloWeave`. From the repository root:

```sh
go test ./...
go vet ./...
go build -o soloweave.exe ./cmd/soloweave
```

`scripts/package-windows.ps1` creates the Windows x64 ZIP and SHA256 checksum under the Git-ignored `dist/` directory. The existing `scripts/build-preview.ps1` remains available for local preview builds. Cross-compilation does not verify runtime behavior on Linux or macOS.

## Verification status

Go tests, vet, build, validation of the five Skills, and a complete Windows CLI smoke test in a temporary project have passed. Linux and macOS binaries have only been cross-compiled. Live handoffs in Claude Code and Cursor remain untested; see the [acceptance procedure](agent-acceptance.md). [GitHub Actions CI](https://github.com/ShiYuandi/SoloWeave/actions/runs/37946714081) passed for commit `b07d057`; check the CI result for later commits separately. Refer to the Releases page for current publication status.

License: [MIT](../../LICENSE) · [Contributing](CONTRIBUTING.md) · [Security](SECURITY.md) · [Changelog](CHANGELOG.md)
