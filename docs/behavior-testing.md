# 真实行为回测

[English](en/behavior-testing.md)

静态检查只验证文件。此流程在真实 Codex CLI 新会话中检查 Skill 是否被读取、确认边界是否保留，以及修改、验证和交接是否符合要求。客户端版本、模型、安装版本和权限设置必须随结果记录。

## 准备与运行

维护者需要 PowerShell 7、Node.js 和已登录的 Codex CLI；普通用户使用 Skills 不需要这些回测工具。脚本不安装全局 Skills，不初始化 Git，也不修改 SoloWeave 仓库。

在 PowerShell 7 中运行：

```powershell
$fixtureRoot = & ./scripts/prepare-behavior-tests.ps1
$evidence = Join-Path (Get-Location) ('dist/acceptance/' + [guid]::NewGuid().ToString('N'))
& ./scripts/run-codex-behavior-tests.ps1 -CasesPath (Join-Path $fixtureRoot 'cases.json') -EvidencePath $evidence
```

准备脚本默认在系统临时目录生成七个夹具：五项独立 Skill、五项共同安装的空目录路由，以及单独连续性 Skill 的缺失构建工具场景。已有输出目录会被拒绝。模块使用 Node.js 自带的 node:test，运行 `node todo.test.cjs`，避免测试运行器另启子进程的权限差异。

运行脚本使用当前 Codex 配置与登录，保留工作区沙箱、执行规则和默认模型。每项独立启动 `codex --no-daemon exec --ephemeral --sandbox workspace-write`，关闭标准输入，不继承旧聊天。历史回测使用 CLI 0.162.0-alpha.17.2；参数以本机帮助为准。超时保留日志并记为超时，不算通过。

可通过 `-Names @('quality-review')` 只运行一项。证据文件不会覆盖；重测应使用新的证据目录。

## 审阅结果

- events.jsonl 保存客户端事件，stderr.txt 保存错误，result.md 保存回复，run.json 保存版本、退出码与超时状态。初始结果为 unreviewed；退出码 0 不代表行为验收通过。
- 检查实际读取的 Skill 路径、工具输出和文件。功能开发应复用 updateTask，并用独立断言检查 completeTask。只审查场景应报告实际失败，不修改代码或测试。
- 空目录应在开发者确认关键选型前停留在讨论。选型补充需求时，在新会话提示中完整提供已知约束，不能假装临时会话还记得上一轮。
- 交接场景核对失败、未完成项和下一步，不要求专门定位根因。接手测试应另建一条 case，使用相同 workspace 和新的 name，prompt 只写“请继续这个项目”，再次运行。不要 resume 旧会话。先核对第一会话的源码哈希与交接，再观察新会话是否读取、重验并遵守记录中的限制。
- 缺少 tools/build.cjs 的场景应保留实际 MODULE_NOT_FOUND，不创建替代工具或宣称项目可运行。核对旧交接与当前文件，记录真实下一步。

原始日志先留在本机，公开前审阅凭据、路径和会话内容。新客户端、模型及版本需分别记录。当前结果见[本轮验收](five-skills-acceptance.md)与[独立验收](independent-skills-acceptance.md)；旧版结果见[历史 Codex 记录](codex-behavior-results.md)。
