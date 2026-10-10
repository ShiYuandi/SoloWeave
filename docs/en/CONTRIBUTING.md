# Contributing

[简体中文](../../CONTRIBUTING.md)

Before a substantial change, read the [development plan](../development-plan.md), [v0.2 design baseline](../design-v0.2.md), and [current handoff](../../.soloweave/context/HANDOFF.md). Then verify the actual code and Git state. The plan and design baseline are maintained in Chinese.

Source development requires Go 1.27.2 or newer. Other users can download the ready-to-run Windows x64 package from [Releases](https://github.com/ShiYuandi/SoloWeave/releases).

- Keep Skills concise and platform-neutral. Put platform differences in installation adapters and rule files.
- When changing CLI behavior, add focused behavioral tests rather than tests that merely repeat the implementation.
- Update `.soloweave/context/STATUS.md` and `CHANGES.md` for important work, and `HANDOFF.md` at handoff.
- For an intentional change to critical approved architecture, review the project contract again, run `soloweave approve`, and retain the new ADR.
- Before committing, run `go test ./...`, `go vet ./...`, and `go build ./...`; report any check you did not run accurately.
- Write Git commit messages in Chinese and briefly describe the change.

Do not commit `.tools/`, `dist/`, secrets, or other machine-generated files. See the [directory guide](directory-structure.md).

## Build from source

Run these commands from the repository root:

```sh
go test ./...
go vet ./...
go build -o soloweave.exe ./cmd/soloweave
```

## Windows download package

Run `./scripts/package-windows.ps1` in Windows PowerShell to create a Windows x64 ZIP and `SHA256SUMS` under the ignored `dist/` directory. The package name uses the version reported by `soloweave.exe version`, so it stays aligned with the executable. Before publishing, extract the ZIP, run its executable, and verify its checksum.

After the code is committed and pushed to the default branch, the repository owner can manually run **Windows draft release** in GitHub Actions. The workflow runs Go tests and vet, packages and checks the executable, then creates a **draft Release** containing the ZIP and checksum file. The owner reviews and publishes the draft on GitHub; running the workflow does not publish it.
