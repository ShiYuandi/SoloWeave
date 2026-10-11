param([string]$Root = (Split-Path -Parent $PSScriptRoot))
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$Root = [IO.Path]::GetFullPath($Root)
$files = @(Get-ChildItem -LiteralPath $Root -File -Filter '*.md')
foreach ($directory in @('docs', 'skills', '.soloweave')) {
    $path = Join-Path $Root $directory
    if (Test-Path -LiteralPath $path) { $files += @(Get-ChildItem -LiteralPath $path -Recurse -File -Filter '*.md') }
}
$missing = @()
foreach ($file in $files) {
    $text = [IO.File]::ReadAllText($file.FullName)
    if ([string]::IsNullOrWhiteSpace($text)) { throw "Empty Markdown document: $($file.FullName)" }
    foreach ($match in [regex]::Matches($text, '\]\((?<target>[^\s)]+)\)')) {
        $target = $match.Groups['target'].Value.Trim('<', '>')
        if ($target -match '^(?:[a-zA-Z][a-zA-Z0-9+.-]*:|#|/)') { continue }
        $target = [Uri]::UnescapeDataString(($target -split '#')[0])
        if (-not $target) { continue }
        $path = [IO.Path]::GetFullPath((Join-Path $file.DirectoryName $target))
        if (-not (Test-Path -LiteralPath $path)) { $missing += "$($file.FullName) -> $target" }
    }
}
if ($missing.Count) { throw ($missing -join "`n") }
Write-Output "PASS relative file links in $($files.Count) Markdown files (web links and heading anchors not checked)"
