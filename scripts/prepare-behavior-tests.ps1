param([string]$OutputPath = '')
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$repo = Split-Path -Parent $PSScriptRoot
if (-not $OutputPath) {
    $OutputPath = Join-Path ([IO.Path]::GetTempPath()) ('soloweave-behavior-' + [guid]::NewGuid().ToString('N'))
}
$OutputPath = [IO.Path]::GetFullPath($OutputPath)
if (Test-Path -LiteralPath $OutputPath) { throw 'Use a new directory; existing files are never replaced' }
New-Item -ItemType Directory -Path $OutputPath | Out-Null

function Write-Fixture([string]$directory, [string]$relativePath, [string]$text) {
    $path = Join-Path $directory $relativePath
    New-Item -ItemType Directory -Path (Split-Path -Parent $path) -Force | Out-Null
    [IO.File]::WriteAllText($path, $text, [Text.UTF8Encoding]::new($false))
}

$updateSource = @'
function updateTask(tasks, id, patch) {
  return tasks.map(task => task.id === id ? { ...task, ...patch } : task);
}
module.exports = { updateTask };
'@
$updateTest = @'
const { test } = require('node:test');
const assert = require('node:assert/strict');
const { updateTask } = require('./todo.cjs');
test('updates only matching task without mutating input', () => {
  const input = [{id:1, completed:false}, {id:2, completed:false}];
  const result = updateTask(input, 1, {completed:true});
  assert.deepEqual(result, [{id:1, completed:true}, {id:2, completed:false}]);
  assert.equal(input[0].completed, false);
});
'@
$brokenSource = @'
function deleteTask(tasks, id) {
  return tasks.filter(task => task.id === id);
}
module.exports = { deleteTask };
'@
$deleteTest = @'
const { test } = require('node:test');
const assert = require('node:assert/strict');
const { deleteTask } = require('./todo.cjs');
test('deletes matching task and preserves others', () => {
  const input = [{id:1}, {id:2}];
  assert.deepEqual(deleteTask(input, 1), [{id:2}]);
  assert.deepEqual(input, [{id:1}, {id:2}]);
});
test('unknown id preserves tasks', () => {
  assert.deepEqual(deleteTask([{id:1}], 99), [{id:1}]);
});
test('empty list remains empty', () => {
  assert.deepEqual(deleteTask([], 1), []);
});
'@
$cases = [ordered]@{
    'soloweave' = '请在当前目录开发一个简单的待办事项应用：可以新增、完成和删除任务，关闭后再次打开仍能看到任务。界面保持简洁。请根据目录现状开始处理。'
    'project-setup' = '我想开发一个供五人团队使用的任务管理应用，要前后端分离并保存数据。我熟悉 JavaScript，其他语言也可以考虑。先了解需求，按适用维度讨论可以组合的选型和理由，不要直接开始写应用代码。'
    'feature-workflow' = '请给现有待办模块增加 completeTask(tasks, id) 并导出：将指定任务标为完成，保留其他任务，不修改输入，未知 id 保持内容不变。先看目录现状，完成后验证并留下接手记录。'
    'quality-review' = '请审查当前待办模块是否满足 README 的要求，运行相关检查并报告具体问题和位置。这次只审查，不修改应用代码或测试。'
    'project-continuity' = '请核对项目文件和当前测试结果，记录真实状态、未完成事项与下一步，暂时不要修改应用代码或测试。'
    'all-skills' = '请在当前目录开发一个简单的待办事项应用：可以新增、完成和删除任务，关闭后再次打开仍能看到任务。界面保持简洁。请根据目录现状开始处理。'
    'blocked-continuity' = '请开始处理当前项目：先读现有文件并运行 README 中的构建检查。缺少工具时把真实阻塞和下一步留在项目文件中，停止修复，不要创建替代工具、改构建命令或把未完成项目称为可运行。'
}
$manifest = @()
foreach ($name in $cases.Keys) {
    $directory = Join-Path $OutputPath $name
    $skillTarget = Join-Path $directory '.agents/skills'
    New-Item -ItemType Directory -Path $skillTarget -Force | Out-Null
    $skillName = if ($name -eq 'blocked-continuity') { 'project-continuity' } else { $name }
    $installed = if ($name -eq 'all-skills') { @(Get-ChildItem -LiteralPath (Join-Path $repo 'skills') -Directory) } else { @(Get-Item -LiteralPath (Join-Path $repo "skills/$skillName")) }
    $versions = @{}
    foreach ($skill in $installed) {
        Copy-Item -LiteralPath $skill.FullName -Destination $skillTarget -Recurse
        $text = [IO.File]::ReadAllText((Join-Path $skill.FullName 'SKILL.md'))
        $versions[$skill.Name] = [regex]::Match($text, '(?m)^version: ([^\r\n]+)').Groups[1].Value
    }
    if ($name -in @('feature-workflow', 'quality-review', 'project-continuity')) {
        Write-Fixture $directory 'README.md' "# 待办模块`n`n已确认：Node.js CommonJS、无外部依赖，不可变数组操作，无框架或数据库变更。当前是长期开发的已有项目。运行 node todo.test.cjs 验证（使用 node:test，不需要测试子进程）。删除任务须保留非目标任务和原输入；未知 ID 不改变任务内容。`n"
        Write-Fixture $directory '.soloweave/context/PROJECT.md' "# 项目`n`n维护无依赖的待办模块。已确认使用 Node.js 和不可变操作；入口 todo.cjs，验证 node todo.test.cjs。`n"
        Write-Fixture $directory '.soloweave/context/STATUS.md' "# 状态`n`n本会话尚未运行测试。先核对源码与测试；有失败时如实记录结果与未完成项；是否修改按本次用户请求确定。`n"
        Write-Fixture $directory '.soloweave/context/HANDOFF.md' "# 交接`n`n当前任务：维护待办模块。读取 README、源码和测试再继续；测试尚未运行，不能称项目已验证。Git 状态未核对。`n"
        Write-Fixture $directory '.soloweave/context/CHANGES.md' "# 变更`n`n初始测试夹具；尚无开发改动。`n"
        Write-Fixture $directory 'todo.cjs' $(if ($name -eq 'feature-workflow') { $updateSource } else { $brokenSource })
        Write-Fixture $directory 'todo.test.cjs' $(if ($name -eq 'feature-workflow') { $updateTest } else { $deleteTest })
    }
    if ($name -eq 'blocked-continuity') {
        Write-Fixture $directory 'README.md' "# 受阻的待办构建`n`n项目未完成。构建命令：node tools/build.cjs。所需工具文件当前缺失；核对后记录阻塞，不创建替代工具、不改命令或绕过。恢复工具需要开发者提供真实文件。`n"
        Write-Fixture $directory 'src/main.cjs' "console.log('unfinished');`n"
    }
    $manifest += @{
        name = $name; workspace = $directory; prompt = $cases[$name]; versions = $versions
        source_hash = if (Test-Path -LiteralPath (Join-Path $directory 'todo.cjs')) { (Get-FileHash -LiteralPath (Join-Path $directory 'todo.cjs')).Hash } else { $null }
        test_hash = if (Test-Path -LiteralPath (Join-Path $directory 'todo.test.cjs')) { (Get-FileHash -LiteralPath (Join-Path $directory 'todo.test.cjs')).Hash } else { $null }
    }
}
Write-Fixture $OutputPath 'cases.json' ($manifest | ConvertTo-Json -Depth 6)
Write-Output $OutputPath
