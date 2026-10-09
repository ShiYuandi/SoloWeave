---
name: quality-review
description: Use when checking a SoloWeave project's code changes, engineering contract, tests, or readiness to hand off or deliver.
---

# Quality review

Inspect the actual change and the relevant project contract. Run the applicable tests and `soloweave check`; use `soloweave doctor` for setup or installation problems. Check for architecture drift, missing handoff details, and newly duplicated code where relevant.

Separate passed, failed, not run, and unavailable checks. Never describe an unrun test as passed. Report concrete findings with affected files and a practical next action. Update the handoff with verified results before an agent switch.
