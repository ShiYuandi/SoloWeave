$ErrorActionPreference = 'Stop'

$root = Split-Path -Parent $PSScriptRoot
$localGo = Join-Path $root '.tools/go/bin/go.exe'
if (Test-Path -LiteralPath $localGo) {
    $go = $localGo
} else {
    $go = (Get-Command go -ErrorAction Stop).Source
}

$dist = Join-Path $root 'dist'
New-Item -ItemType Directory -Force -Path $dist | Out-Null
$exe = Join-Path $dist 'soloweave.exe'

Push-Location $root
try {
    $env:CGO_ENABLED = '0'
    $env:GOOS = 'windows'
    $env:GOARCH = 'amd64'
    & $go build -trimpath -o $exe ./cmd/soloweave
    if ($LASTEXITCODE -ne 0) { throw 'Windows package build failed.' }

    $version = (& $exe version | Out-String).Trim()
    if ($LASTEXITCODE -ne 0 -or $version -notmatch '^[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9A-Za-z.-]+)?$') {
        throw "Invalid version reported by soloweave.exe: $version"
    }

    $license = Join-Path $dist 'LICENSE'
    $readmeZh = Join-Path $dist 'README.zh-CN.md'
    $readmeEn = Join-Path $dist 'README.en.md'
    Copy-Item -LiteralPath (Join-Path $root 'LICENSE') -Destination $license -Force
    $utf8 = [System.Text.UTF8Encoding]::new($false)
    $zhTemplate = [System.IO.File]::ReadAllText((Join-Path $root 'docs/package-readme.zh-CN.md'))
    $enTemplate = [System.IO.File]::ReadAllText((Join-Path $root 'docs/package-readme.en.md'))
    foreach ($template in @($zhTemplate, $enTemplate)) {
        if (-not $template.Contains('{{VERSION}}')) { throw 'Package README version placeholder missing.' }
    }
    [System.IO.File]::WriteAllText($readmeZh, $zhTemplate.Replace('{{VERSION}}', $version), $utf8)
    [System.IO.File]::WriteAllText($readmeEn, $enTemplate.Replace('{{VERSION}}', $version), $utf8)

    $archiveName = "soloweave-v$version-windows-amd64.zip"
    $archive = Join-Path $dist $archiveName
    Compress-Archive -LiteralPath @($exe, $license, $readmeZh, $readmeEn) -DestinationPath $archive -Force

    $hash = (Get-FileHash -Algorithm SHA256 -LiteralPath $archive).Hash.ToLowerInvariant()
    [System.IO.File]::WriteAllText(
        (Join-Path $dist 'SHA256SUMS'),
        "$hash  $archiveName`n",
        [System.Text.UTF8Encoding]::new($false)
    )
    Write-Output $archive
    Write-Output (Join-Path $dist 'SHA256SUMS')
} finally {
    Pop-Location
}
