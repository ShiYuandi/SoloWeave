Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
Add-Type -AssemblyName System.IO.Compression

$root = Split-Path -Parent $PSScriptRoot
$sourceRoot = Join-Path $root 'skills'
$outputRoot = Join-Path $root 'dist/skillhub/skills-first'
New-Item -ItemType Directory -Path $outputRoot -Force | Out-Null

$checksums = @()
$slugs = @{}
$skills = @(Get-ChildItem -LiteralPath $sourceRoot -Directory | Sort-Object Name)
$expectedNames = @('feature-workflow', 'project-continuity', 'project-setup', 'quality-review', 'soloweave', 'systematic-debugging')
if (@(Compare-Object -ReferenceObject $expectedNames -DifferenceObject @($skills.Name)).Count -ne 0) {
    throw "Unexpected SoloWeave skills. Expected: $($expectedNames -join ', '); found: $($skills.Name -join ', ')."
}

foreach ($skill in $skills) {
    $skillFile = Join-Path $skill.FullName 'SKILL.md'
    $content = Get-Content -LiteralPath $skillFile -Raw -Encoding UTF8
    $match = [regex]::Match($content, '\A---\r?\n(?<frontmatter>.*?)\r?\n---\r?\n', 'Singleline')
    if (-not $match.Success) { throw "Invalid SKILL.md frontmatter: $skillFile" }
    $frontmatter = $match.Groups['frontmatter'].Value
    $values = @{}
    foreach ($key in @('name', 'description', 'slug', 'version', 'displayName', 'summary', 'license')) {
        $field = [regex]::Match($frontmatter, "(?m)^$([regex]::Escape($key)):\s*(?<value>\S.*)$")
        if (-not $field.Success) { throw "Missing $key in $skillFile" }
        $values[$key] = $field.Groups['value'].Value.Trim()
    }
    if ($values['name'] -cne $skill.Name) { throw "Skill name does not match directory: $skillFile" }
    if ($values['name'] -cnotmatch '^[a-z0-9]+(?:-[a-z0-9]+)*$' -or $values['name'].Length -gt 64) {
        throw "Invalid Agent Skills name: $($values['name'])"
    }
    if ($values['description'].Length -gt 1024) { throw "Skill description too long: $skillFile" }
    if ($values['slug'] -cnotmatch '^[a-z0-9]+(?:-[a-z0-9]+)*$' -or $values['slug'].Length -gt 128) {
        throw "Invalid SkillHub slug: $($values['slug'])"
    }
    if ($values['version'] -cnotmatch '^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$') {
        throw "SkillHub requires a numeric X.Y.Z version: $($values['version'])"
    }
    if ($slugs.ContainsKey($values['slug'])) { throw "Duplicate SkillHub slug: $($values['slug'])" }
    $slugs[$values['slug']] = $true

    $zipName = "$($values['slug'])-$($values['version']).zip"
    $zipPath = Join-Path $outputRoot $zipName
    $zipStream = [System.IO.File]::Open($zipPath, [System.IO.FileMode]::Create)
    try {
        $archive = [System.IO.Compression.ZipArchive]::new($zipStream, [System.IO.Compression.ZipArchiveMode]::Create)
        try {
            $entry = $archive.CreateEntry("$($skill.Name)/SKILL.md")
            $entryStream = $entry.Open()
            try {
                $bytes = [System.IO.File]::ReadAllBytes($skillFile)
                $entryStream.Write($bytes, 0, $bytes.Length)
            } finally {
                $entryStream.Dispose()
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
