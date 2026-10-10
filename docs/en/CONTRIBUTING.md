# Contributing

[简体中文](../../CONTRIBUTING.md)

Before a substantial change, read the [Skills-first design](../skills-first-design.md), [migration plan](../skills-first-plan.md), and [current handoff](../../.soloweave/context/HANDOFF.md). Then verify the actual code and Git state. The earlier CLI scope remains in the [development plan](../development-plan.md) and [v0.2 baseline](../design-v0.2.md), maintained in Chinese.

Editing the new root `skills/` does not require Go. Only maintenance of the legacy CLI requires Go 1.27.2 or newer. Ordinary users can install Skills without downloading or building the program.

- Edit new Skills only under root `skills/`; `internal/bundle/assets/skills/` is the legacy CLI snapshot. Keep Skills concise and platform-neutral.
- When changing CLI behavior, add focused behavioral tests rather than tests that merely repeat the implementation.
- Update `.soloweave/context/STATUS.md` and `CHANGES.md` for important work, and `HANDOFF.md` at handoff.
- For an intentional change to critical approved architecture, obtain explicit developer approval and add an ADR. Changes to the legacy CLI's own contract still follow its `approve` flow.
- When editing a new Skill, run `scripts/package-skillhub.ps1` and inspect the ZIPs. Run `go test ./...`, `go vet ./...`, and `go build ./...` when changing the legacy CLI. Accurately report checks you did not run.
- Write Git commit messages in Chinese and briefly describe the change.

Do not commit `.tools/`, `dist/`, secrets, or other machine-generated files. See the [directory guide](directory-structure.md).

## Build from source

These commands are for maintaining the legacy Go CLI. Run them from the repository root:

```sh
go test ./...
go vet ./...
go build -o soloweave.exe ./cmd/soloweave
```

## Windows download package

Run `./scripts/package-windows.ps1` in Windows PowerShell to create a Windows x64 ZIP and `SHA256SUMS` under the ignored `dist/` directory. The package name uses the version reported by `soloweave.exe version`, so it stays aligned with the executable. Before publishing, extract the ZIP, run its executable, and verify its checksum.

After the code is committed and pushed to the default branch, the repository owner can manually run **Windows draft release** in GitHub Actions. The workflow runs Go tests and vet, packages and checks the executable, then creates a **draft Release** containing the ZIP and checksum file. The owner reviews and publishes the draft on GitHub; running the workflow does not publish it.
