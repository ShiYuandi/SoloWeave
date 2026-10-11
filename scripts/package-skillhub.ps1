param(
    [string]$Root = (Split-Path -Parent $PSScriptRoot),
    [string]$OutputPath = ''
)
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
Add-Type -AssemblyName System.IO.Compression

$Root = [IO.Path]::GetFullPath($Root)
& (Join-Path $PSScriptRoot 'check-skills.ps1') -Root $Root
$sourceRoot = Join-Path $Root 'skills'
if (-not $OutputPath) { $OutputPath = Join-Path $Root 'dist/skills' }
$outputRoot = [IO.Path]::GetFullPath($OutputPath)
$sourcePrefix = $sourceRoot.TrimEnd([IO.Path]::DirectorySeparatorChar) + [IO.Path]::DirectorySeparatorChar
if ($outputRoot -eq $sourceRoot -or $outputRoot.StartsWith($sourcePrefix, [StringComparison]::OrdinalIgnoreCase)) { throw 'Package output must be outside skills/' }
New-Item -ItemType Directory -Path $outputRoot -Force | Out-Null

$checksums = @()
$skills = @(Get-ChildItem -LiteralPath $sourceRoot -Directory | Sort-Object Name)

foreach ($skill in $skills) {
    $skillFile = Join-Path $skill.FullName 'SKILL.md'
    $content = Get-Content -LiteralPath $skillFile -Raw -Encoding UTF8
    $match = [regex]::Match($content, '\A---\r?\n(?<frontmatter>.*?)\r?\n---\r?\n', 'Singleline')
    $frontmatter = $match.Groups['frontmatter'].Value
    $values = @{}
    foreach ($key in @('slug', 'version')) {
        $field = [regex]::Match($frontmatter, "(?m)^$([regex]::Escape($key)):[ \t]*(?<value>\S.*)$")
        $values[$key] = $field.Groups['value'].Value.Trim().Trim('"').Trim("'")
    }

    $zipName = "$($values['slug'])-$($values['version']).zip"
    $zipPath = Join-Path $outputRoot $zipName
    $zipStream = [System.IO.File]::Open($zipPath, [System.IO.FileMode]::Create)
    try {
        $archive = [System.IO.Compression.ZipArchive]::new($zipStream, [System.IO.Compression.ZipArchiveMode]::Create)
        try {
            $sourceFiles = @(Get-ChildItem -LiteralPath $skill.FullName -File -Recurse -Force | Sort-Object FullName)
            foreach ($sourceFile in $sourceFiles) {
                $relativePath = [System.IO.Path]::GetRelativePath($skill.FullName, $sourceFile.FullName).Replace('\', '/')
                $entry = $archive.CreateEntry("$($skill.Name)/$relativePath")
                $entry.LastWriteTime = [DateTimeOffset]::new(2000, 1, 1, 0, 0, 0, [TimeSpan]::Zero)
                $entryStream = $entry.Open()
                try {
                    $bytes = [System.IO.File]::ReadAllBytes($sourceFile.FullName)
                    $entryStream.Write($bytes, 0, $bytes.Length)
                } finally {
                    $entryStream.Dispose()
                }
            }
        } finally {
            $archive.Dispose()
        }
    } finally {
        $zipStream.Dispose()
    }
    $hash = (Get-FileHash -LiteralPath $zipPath -Algorithm SHA256).Hash.ToLowerInvariant()
    $checksums += "$hash  $zipName"
    Write-Output "$($values['slug'])@$($values['version']): $zipPath"
}

$checksums | Set-Content -LiteralPath (Join-Path $outputRoot 'SHA256SUMS') -Encoding ascii
