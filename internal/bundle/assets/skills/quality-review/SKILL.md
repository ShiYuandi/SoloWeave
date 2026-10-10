---
name: quality-review
description: Use when checking a SoloWeave project's code changes, engineering contract, tests, or readiness to hand off or deliver.
slug: shiyuandi-soloweave-quality-review
version: 0.1.0
displayName: SoloWeave 质量检查
summary: 核对代码、配置与真实测试结果，并明确报告失败或未执行的检查。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# Quality review

Inspect the actual change and the relevant project contract. Run the applicable tests and `soloweave check`; use `soloweave doctor` for setup or installation problems. Check for architecture drift, missing handoff details, and newly duplicated code where relevant.

Separate passed, failed, not run, and unavailable checks. Never describe an unrun test as passed. Report concrete findings with affected files and a practical next action. Update the handoff with verified results before an agent switch.
