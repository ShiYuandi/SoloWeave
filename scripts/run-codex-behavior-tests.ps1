param(
    [Parameter(Mandatory)][string]$CasesPath,
    [Parameter(Mandatory)][string]$EvidencePath,
    [string[]]$Names = @('soloweave', 'project-setup', 'feature-workflow', 'quality-review', 'project-continuity', 'all-skills', 'blocked-continuity'),
    [int]$TimeoutSeconds = 240
)
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
if ($TimeoutSeconds -lt 1) { throw 'TimeoutSeconds must be positive' }
$cases = @(Get-Content -LiteralPath $CasesPath -Raw | ConvertFrom-Json)
$EvidencePath = [IO.Path]::GetFullPath($EvidencePath)
New-Item -ItemType Directory -Path $EvidencePath -Force | Out-Null
$codex = (Get-Command codex -ErrorAction Stop).Source
foreach ($name in $Names) {
    $selected = @($cases | Where-Object name -CEQ $name)
    if ($selected.Count -ne 1) { throw "Expected exactly one case: $name" }
    $fixture = $selected[0]
    $prefix = Join-Path $EvidencePath $name
    foreach ($suffix in @('.events.jsonl', '.stderr.txt', '.result.md', '.run.json')) {
        if (Test-Path -LiteralPath ($prefix + $suffix)) { throw "Existing evidence will not be overwritten: $prefix$suffix" }
    }
    $start = [Diagnostics.ProcessStartInfo]::new($codex)
    $start.UseShellExecute = $false
    $start.CreateNoWindow = $true
    $start.RedirectStandardInput = $true
    $start.RedirectStandardOutput = $true
    $start.RedirectStandardError = $true
    foreach ($argument in @('--no-daemon', 'exec', '--ephemeral', '--sandbox', 'workspace-write', '--skip-git-repo-check', '--json', '--cd', $fixture.workspace, '--output-last-message', ($prefix + '.result.md'), $fixture.prompt)) {
        $start.ArgumentList.Add($argument)
    }
    $process = [Diagnostics.Process]::new()
    $process.StartInfo = $start
    $began = [DateTimeOffset]::Now
    Write-Output "START $name"
    try {
        if (-not $process.Start()) { throw "Cannot start Codex for $name" }
        $process.StandardInput.Close()
        $stdout = $process.StandardOutput.ReadToEndAsync()
        $stderr = $process.StandardError.ReadToEndAsync()
        $timedOut = -not $process.WaitForExit($TimeoutSeconds * 1000)
        if ($timedOut) { $process.Kill($true); $process.WaitForExit() }
        [IO.File]::WriteAllText(($prefix + '.events.jsonl'), $stdout.GetAwaiter().GetResult(), [Text.UTF8Encoding]::new($false))
        [IO.File]::WriteAllText(($prefix + '.stderr.txt'), $stderr.GetAwaiter().GetResult(), [Text.UTF8Encoding]::new($false))
        $record = @{
            name = $name; started = $began.ToString('o'); finished = [DateTimeOffset]::Now.ToString('o')
            workspace = $fixture.workspace; versions = $fixture.versions
            exit_code = $process.ExitCode; timed_out = $timedOut
            acceptance = 'unreviewed'; note = 'Process exit is not behavioral acceptance; inspect tool events and files.'
        }
        [IO.File]::WriteAllText(($prefix + '.run.json'), ($record | ConvertTo-Json -Depth 5), [Text.UTF8Encoding]::new($false))
        Write-Output "FINISH $name exit=$($process.ExitCode) timed_out=$timedOut (behavior requires review)"
    } finally { $process.Dispose() }
}
