# Live behavior testing

[简体中文](../behavior-testing.md)

Static checks validate files. These fixtures use real, fresh Codex CLI sessions to check Skill reading, confirmation, implementation, validation, and handoff. Record client version, configured model, installed versions, and permissions separately.

Maintainers need PowerShell 7, Node.js, and an authenticated Codex CLI. Users do not need these testing tools. The scripts do not install global Skills, initialize Git, or modify SoloWeave.

```powershell
$fixtureRoot = & ./scripts/prepare-behavior-tests.ps1
$evidence = Join-Path (Get-Location) ('dist/acceptance/' + [guid]::NewGuid().ToString('N'))
& ./scripts/run-codex-behavior-tests.ps1 -CasesPath (Join-Path $fixtureRoot 'cases.json') -EvidencePath $evidence
```

Preparation creates eight fresh temporary fixtures: six standalone Skills, all-six new-project routing, and missing-build-tool continuity. Existing directories are rejected. Module tests run Node's built-in node:test through `node todo.test.cjs`, avoiding test-worker process permissions.

Each case starts `codex --no-daemon exec --ephemeral --sandbox workspace-write` with closed stdin. Current configuration, execution rules, authentication, and the default model are preserved. This round used CLI 0.162.0-alpha.17.2; consult local help for options. Timeout is recorded, not accepted as a pass. Use `-Names @('quality-review')` for one case and a new evidence directory for each retest.

Artifacts include events, stderr, final responses, and run records initially marked unreviewed. Exit zero is not acceptance: inspect actual Skill paths, tool output, files, and independent assertions. Feature work should reuse updateTask; debugging must reproduce failures and preserve original tests; review-only cases must leave source unchanged. New projects must await major choices.

For resumption, add a case with a new name, the same workspace, and only “请继续这个项目” as the prompt. Start a fresh session, not a resume command. Verify checkpoint files, source hashes, actual checks, and prior developer constraints. For selection follow-ups, explicitly provide all known requirements rather than assuming chat memory.

The blocked fixture must preserve MODULE_NOT_FOUND for tools/build.cjs without fabricating a replacement or claiming the project runs. Keep raw logs local until reviewed. Record each client/model/version independently. See [acceptance](../independent-skills-acceptance.md) and [Codex results](../codex-retest.md).
