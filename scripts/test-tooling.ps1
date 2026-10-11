Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$repo = Split-Path -Parent $PSScriptRoot
$check = Join-Path $PSScriptRoot 'check-skills.ps1'
$package = Join-Path $PSScriptRoot 'package-skillhub.ps1'
$scratch = Join-Path ([IO.Path]::GetTempPath()) ('soloweave-tooling-' + [guid]::NewGuid().ToString('N'))
$passed = 0

function New-Fixture {
    $fixture = Join-Path $scratch ([guid]::NewGuid().ToString('N'))
    New-Item -ItemType Directory -Path $fixture -Force | Out-Null
    Copy-Item -LiteralPath (Join-Path $repo 'skills') -Destination $fixture -Recurse
    return $fixture
}

function Expect-Rejection([string]$label, [scriptblock]$mutate) {
    $fixture = New-Fixture
    & $mutate $fixture
    $output = & pwsh -NoProfile -File $check -Root $fixture 2>&1
    if ($LASTEXITCODE -eq 0) { throw "Accepted invalid fixture: $label" }
    Write-Output "PASS rejects $label"
    $script:passed++
}

try {
    New-Item -ItemType Directory -Path $scratch -Force | Out-Null
    & $check -Root $repo
    if (-not $?) { throw 'Source validation failed' }
    $passed++
    Expect-Rejection 'missing description' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        $text = [IO.File]::ReadAllText($path) -replace '(?m)^description:.*\r?\n', ''
        [IO.File]::WriteAllText($path, $text)
    }
    Expect-Rejection 'invalid version' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        [IO.File]::WriteAllText($path, ([IO.File]::ReadAllText($path) -replace '(?m)^version:.*$', 'version: latest'))
    }
    Expect-Rejection 'empty description' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        [IO.File]::WriteAllText($path, ([IO.File]::ReadAllText($path) -replace '(?m)^description:.*$', 'description:'))
    }
    Expect-Rejection 'duplicate slug' { param($f)
        $path = Join-Path $f 'skills/project-setup/SKILL.md'
        [IO.File]::WriteAllText($path, ([IO.File]::ReadAllText($path) -replace '(?m)^slug:.*$', 'slug: shiyuandi-soloweave'))
    }
    Expect-Rejection 'missing reference' { param($f)
        Remove-Item -LiteralPath (Join-Path $f 'skills/project-setup/references/selection-guide.md')
    }
    Expect-Rejection 'reference escapes skill' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        Add-Content -LiteralPath $path -Value '[outside](../project-setup/SKILL.md)'
    }
    Expect-Rejection 'missing skill' { param($f)
        Remove-Item -LiteralPath (Join-Path $f 'skills/quality-review/SKILL.md')
    }
    Expect-Rejection 'duplicate field' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        [IO.File]::WriteAllText($path, ([IO.File]::ReadAllText($path) -replace '(?m)^name: soloweave', "name: soloweave`nname: soloweave"))
    }
    Expect-Rejection 'duplicate empty field' { param($f)
        $path = Join-Path $f 'skills/soloweave/SKILL.md'
        [IO.File]::WriteAllText($path, ([IO.File]::ReadAllText($path) -replace '(?m)^name: soloweave', "name: soloweave`nname:"))
    }
    foreach ($case in @('empty document', 'missing document link')) {
        $fixture = New-Fixture
        $text = if ($case -eq 'empty document') { '' } else { '[missing](not-present.md)' }
        [IO.File]::WriteAllText((Join-Path $fixture 'README.md'), $text)
        $output = & pwsh -NoProfile -File (Join-Path $PSScriptRoot 'check-docs.ps1') -Root $fixture 2>&1
        if ($LASTEXITCODE -eq 0) { throw "Accepted invalid fixture: $case" }
        $passed++
        Write-Output "PASS rejects $case"
    }
    $outputDir = Join-Path $scratch 'packages'
    $fixture = New-Fixture
    $output = & pwsh -NoProfile -File $package -Root $fixture -OutputPath (Join-Path $fixture 'skills/output') 2>&1
    if ($LASTEXITCODE -eq 0) { throw 'Accepted package output inside source' }
    $passed++
    Write-Output 'PASS rejects package output inside source'
    & $package -Root $repo -OutputPath $outputDir
    $firstManifest = [IO.File]::ReadAllText((Join-Path $outputDir 'SHA256SUMS'))
    Add-Type -AssemblyName System.IO.Compression
    $lines = @(Get-Content -LiteralPath (Join-Path $outputDir 'SHA256SUMS'))
    if ($lines.Count -ne 5) { throw 'Expected five checksum entries' }
    foreach ($line in $lines) {
        $expectedHash, $filename = $line -split '  ', 2
        $zipPath = Join-Path $outputDir $filename
        if ((Get-FileHash -LiteralPath $zipPath -Algorithm SHA256).Hash.ToLowerInvariant() -ne $expectedHash) { throw "Checksum mismatch: $filename" }
        $zip = [IO.Compression.ZipFile]::OpenRead($zipPath)
        try {
            $skillName = ($zip.Entries[0].FullName -split '/')[0]
            $sources = @(Get-ChildItem -LiteralPath (Join-Path $repo "skills/$skillName") -File -Recurse -Force)
            if ($sources.Count -ne $zip.Entries.Count) { throw "Package file count mismatch: $filename" }
            foreach ($entry in $zip.Entries) {
                $source = Join-Path $repo ('skills/' + $entry.FullName)
                $stream = $entry.Open()
                try { $actualHash = [Convert]::ToHexString([Security.Cryptography.SHA256]::HashData($stream)) }
                finally { $stream.Dispose() }
                if ($actualHash -ne (Get-FileHash -LiteralPath $source -Algorithm SHA256).Hash) { throw "Package bytes differ: $($entry.FullName)" }
            }
        } finally { $zip.Dispose() }
    }
    $passed++
    Write-Output 'PASS five package checksums and every resource match source'
    & $package -Root $repo -OutputPath $outputDir
    if ([IO.File]::ReadAllText((Join-Path $outputDir 'SHA256SUMS')) -ne $firstManifest) { throw 'Repeated packaging changed checksums' }
    $passed++
    Write-Output "PASS deterministic packaging; $passed checks completed"
} finally {
    $resolvedScratch = [IO.Path]::GetFullPath($scratch)
    $tempPrefix = [IO.Path]::GetFullPath([IO.Path]::GetTempPath()).TrimEnd([IO.Path]::DirectorySeparatorChar) + [IO.Path]::DirectorySeparatorChar
    if ($resolvedScratch.StartsWith($tempPrefix, [StringComparison]::OrdinalIgnoreCase) -and (Split-Path $resolvedScratch -Leaf).StartsWith('soloweave-tooling-')) {
        if (Test-Path -LiteralPath $resolvedScratch) { Remove-Item -LiteralPath $resolvedScratch -Recurse -Force }
    }
}

# Expected failing child processes leave LASTEXITCODE at 1. Report the suite's
# successful result explicitly so CI wrappers do not reuse that negative case.
exit 0
