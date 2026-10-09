---
name: soloweave
description: Use when starting or resuming work in a SoloWeave project and the task needs routing to project setup, feature development, continuity, or quality review.
---

# SoloWeave

Read `.soloweave/project.yaml` for approved engineering decisions and `.soloweave/context/HANDOFF.md` for the current task. Check relevant source files and Git state before trusting the handoff. If no project contract exists, use `project-setup`. For a feature or fix, use `feature-workflow`; for a pause or agent switch, use `project-continuity`; for verification, use `quality-review`.

The developer chooses major architecture changes. Do not treat a draft configuration as approved. Keep the response proportional to the task; a small edit need not invoke the full workflow.
