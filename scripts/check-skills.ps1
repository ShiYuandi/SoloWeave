param([string]$Root = (Split-Path -Parent $PSScriptRoot))
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$sourceRoot = Join-Path ([IO.Path]::GetFullPath($Root)) 'skills'
$expected = @('feature-workflow', 'project-continuity', 'project-setup', 'quality-review', 'soloweave', 'systematic-debugging')
$skills = @(Get-ChildItem -LiteralPath $sourceRoot -Directory -Force | Sort-Object Name)
if (@(Compare-Object $expected @($skills.Name)).Count -ne 0) { throw 'Expected exactly the six SoloWeave Skill directories' }
$slugs = @{}
foreach ($skill in $skills) {
    $items = @($skill) + @(Get-ChildItem -LiteralPath $skill.FullName -Recurse -Force)
    if (@($items | Where-Object { $_.Attributes -band [IO.FileAttributes]::ReparsePoint }).Count) { throw "Linked resources are not allowed in packages: $($skill.Name)" }
    $skillFile = Join-Path $skill.FullName 'SKILL.md'
    $content = [IO.File]::ReadAllText($skillFile)
    $header = [regex]::Match($content, '\A---\r?\n(?<fields>.*?)\r?\n---\r?\n', 'Singleline')
    if (-not $header.Success) { throw "Invalid frontmatter: $skillFile" }
    $values = @{}
    foreach ($key in @('name', 'description', 'slug', 'version', 'displayName', 'summary', 'license')) {
        $matches = [regex]::Matches($header.Groups['fields'].Value, "(?m)^$([regex]::Escape($key)):[ \t]*(?<value>[^\r\n]*)\r?$")
        if ($matches.Count -ne 1) { throw "Expected one $key field: $skillFile" }
        $value = $matches[0].Groups['value'].Value.Trim().Trim('"').Trim("'")
        if ([string]::IsNullOrWhiteSpace($value) -or $value -in @('|', '>')) { throw "Expected a single-line scalar for $key in $skillFile" }
        $values[$key] = $value
    }
    if ($values.name -cne $skill.Name -or $values.name -cnotmatch '^[a-z0-9]+(?:-[a-z0-9]+)*$' -or $values.name.Length -gt 64) { throw "Invalid Skill name: $skillFile" }
    if ($values.description.Length -gt 1024 -or $values.description -match '[<>]') { throw "Invalid description: $skillFile" }
    if ($values.slug -cnotmatch '^[a-z0-9]+(?:-[a-z0-9]+)*$' -or $values.slug.Length -gt 128 -or $slugs.ContainsKey($values.slug)) { throw "Invalid or duplicate slug: $skillFile" }
    $slugs[$values.slug] = $true
    if ($values.version -cnotmatch '^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$') { throw "Expected numeric X.Y.Z version: $skillFile" }
    if ($values.license -ne 'MIT') { throw "Unexpected license: $skillFile" }
    $prefix = $skill.FullName.TrimEnd([IO.Path]::DirectorySeparatorChar) + [IO.Path]::DirectorySeparatorChar
    foreach ($markdown in @(Get-ChildItem -LiteralPath $skill.FullName -Recurse -File -Filter '*.md')) {
        $text = [IO.File]::ReadAllText($markdown.FullName)
        foreach ($link in [regex]::Matches($text, '\]\((?<target>[^\s)]+)\)')) {
            $target = $link.Groups['target'].Value
            if ($target -match '^(?:[a-zA-Z][a-zA-Z0-9+.-]*:|#)') { continue }
            $target = [Uri]::UnescapeDataString(($target -split '#')[0])
            $path = [IO.Path]::GetFullPath((Join-Path $markdown.DirectoryName $target))
            if (-not $path.StartsWith($prefix, [StringComparison]::OrdinalIgnoreCase) -or -not (Test-Path -LiteralPath $path)) { throw "Missing or escaping resource: $($markdown.FullName) -> $target" }
        }
    }
}
Write-Output 'PASS six Skills: metadata, unique slugs, versions, resources and package boundaries'
