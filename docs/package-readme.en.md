# SoloWeave v{{VERSION}} Windows x64 package

This ZIP contains a ready-to-run `soloweave.exe`, Chinese and English instructions, and the MIT license. **You do not need Go or a source build.**

## Get started

1. Extract the ZIP and add its directory to the current terminal's `PATH`, or use the full path to `soloweave.exe` in the commands below.
2. In your target project directory, run:

```powershell
soloweave init
```

Review the generated `.soloweave/project.yaml`. After the developer approves the technical choices, run:

```powershell
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave context checkpoint --task "project setup" --summary "configuration approved and Skills installed" --next "start development"
soloweave check
```

List only the AI platforms you use in `--agents`. `install --dry-run` previews changes; a regular `install` writes the five Skills and project rules. `check` needs a checkpoint first. Use `soloweave context checkpoint` again for handoffs and `soloweave context resume` in a new session. Omit `--verification` when no test ran.

If you only want the development conventions, you can install the five Skills from SkillHub without the CLI. See the [GitHub README](https://github.com/ShiYuandi/SoloWeave/blob/master/docs/en/README.md) and [Skills-only guide](https://github.com/ShiYuandi/SoloWeave/blob/master/docs/en/skill-only-workflow.md) for both paths and all commands. This file is generated for the package's version; check [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) for the latest public build.
