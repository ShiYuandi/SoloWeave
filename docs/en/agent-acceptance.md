# Cross-agent handoff acceptance

[简体中文（主版本）](../agent-acceptance.md)

Claude Code and Cursor are not installed on the current development machine. Automated tests cover generated files and checkpoint behavior, but a live handoff in those clients remains unverified.

When the clients are available, run this manual acceptance test:

1. Create a disposable project and run `soloweave init`, `soloweave approve`, and `soloweave install --agents codex,claude,cursor`.
2. Start a small feature in Agent A. Update `STATUS.md` and `CHANGES.md` to reflect actual progress. Run a real check and save a checkpoint with its exact result.
3. Close Agent A. Open the same project in Agent B without the previous chat history and ask it to continue using only project files.
4. Confirm Agent B can identify the goal, approved stack, completed and remaining work, recent changes, actual verification, and next step. Modify a source file after the checkpoint and confirm Agent B recognizes that the handoff may be stale.
5. Have Agent B finish the task, run checks, and update the status and handoff. Confirm it does not change critical architecture without developer approval and a new ADR.

Record client versions and outcomes. Until this test is complete, do not claim live Claude Code or Cursor compatibility has been verified.
