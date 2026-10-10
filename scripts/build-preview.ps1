$ErrorActionPreference = 'Stop'
$root = Split-Path -Parent $PSScriptRoot
$localGo = Join-Path $root '.tools/go/bin/go.exe'
if (Test-Path -LiteralPath $localGo) { $go = $localGo } else { $go = (Get-Command go -ErrorAction Stop).Source }
$dist = Join-Path $root 'dist'
New-Item -ItemType Directory -Force -Path $dist | Out-Null
$targets = @(
    @{ OS='windows'; Arch='amd64'; Ext='.exe' },
    @{ OS='linux'; Arch='amd64'; Ext='' },
    @{ OS='darwin'; Arch='amd64'; Ext='' },
    @{ OS='darwin'; Arch='arm64'; Ext='' }
)
Push-Location $root
try {
    $version = (& $go run ./cmd/soloweave version | Out-String).Trim()
    if ($LASTEXITCODE -ne 0 -or $version -notmatch '^[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9A-Za-z.-]+)?$') {
        throw "Invalid version reported by SoloWeave: $version"
    }
    $env:CGO_ENABLED = '0'
    foreach ($target in $targets) {
        $env:GOOS = $target.OS
        $env:GOARCH = $target.Arch
        $output = Join-Path $dist ("soloweave-v{0}-{1}-{2}{3}" -f $version, $target.OS, $target.Arch, $target.Ext)
        & $go build -o $output ./cmd/soloweave
        if ($LASTEXITCODE -ne 0) { throw "Build failed: $($target.OS)/$($target.Arch)" }
        Write-Output $output
    }
    Get-ChildItem -LiteralPath $dist -File | Where-Object Name -ne 'SHA256SUMS' | Sort-Object Name | ForEach-Object {
        '{0}  {1}' -f (Get-FileHash -Algorithm SHA256 -LiteralPath $_.FullName).Hash.ToLowerInvariant(), $_.Name
    } | Set-Content -LiteralPath (Join-Path $dist 'SHA256SUMS') -Encoding utf8
} finally {
    Pop-Location
}
