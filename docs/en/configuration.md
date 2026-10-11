# Project records and decisions

[简体中文](../configuration.md)

Current SoloWeave uses Skills and Markdown records. No initialization command, YAML contract, or dedicated executable is required.

Reuse existing project documentation first. For continuing work, default records are `PROJECT.md`, `STATUS.md`, `HANDOFF.md`, and `CHANGES.md` under `.soloweave/context/`, with decisions in `.soloweave/decisions/ADR-NNNN.md`.

Keep proposed technologies and architecture pending until the developer confirms them. Record alternatives, reasons, constraints, and confirmation in an ADR; significant later changes supersede earlier decisions with a new ADR.

Skills do not supply machine approval digests or automatic stale-check detection. Verify records against actual files and Git. See [continuity](project-continuity.md), [legacy migration](cli-migration.md), and the [archived YAML guide](../archive/go-cli/en/configuration.md).
