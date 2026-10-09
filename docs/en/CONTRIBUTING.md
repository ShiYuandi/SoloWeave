# Contributing

[简体中文](../../CONTRIBUTING.md)

Before a substantial change, read the [development plan](../development-plan.md), [v0.2 design baseline](../design-v0.2.md), and [current handoff](../../.soloweave/context/HANDOFF.md). Then verify the actual code and Git state. The plan and design baseline are maintained in Chinese.

- Keep Skills concise and platform-neutral. Put platform differences in installation adapters and rule files.
- When changing CLI behavior, add focused behavioral tests rather than tests that merely repeat the implementation.
- Update `.soloweave/context/STATUS.md` and `CHANGES.md` for important work, and `HANDOFF.md` at handoff.
- For an intentional change to critical approved architecture, review the project contract again, run `soloweave approve`, and retain the new ADR.
- Before committing, run `go test ./...`, `go vet ./...`, and `go build ./...`; report any check you did not run accurately.
- Write Git commit messages in Chinese and briefly describe the change.

Do not commit `.tools/`, `dist/`, secrets, or other machine-generated files. See the [directory guide](directory-structure.md).
