param(
    [string]$InstallDir = "$HOME\bin"
)

$ErrorActionPreference = "Stop"

$repo = "Pointdexter37/Kin"
$apiUri = "https://api.github.com/repos/$repo/releases/latest"
$headers = @{ "User-Agent" = "kin-installer" }

Write-Host "Finding the latest Kin release..."
$release = Invoke-RestMethod -Uri $apiUri -Headers $headers
$asset = $release.assets | Where-Object { $_.name -match "^Kin_[^_]+_windows_amd64\.zip$" } | Select-Object -First 1
$checksumAsset = $release.assets | Where-Object { $_.name -eq "checksums.txt" } | Select-Object -First 1

if ($null -eq $asset -or $null -eq $checksumAsset) {
    throw "The latest release does not contain the expected Windows archive and checksums."
}

$tempDir = Join-Path ([System.IO.Path]::GetTempPath()) ("kin-install-" + [guid]::NewGuid())
New-Item -ItemType Directory -Path $tempDir | Out-Null

try {
    $archivePath = Join-Path $tempDir $asset.name
    $checksumPath = Join-Path $tempDir "checksums.txt"
    $extractDir = Join-Path $tempDir "extracted"

    Invoke-WebRequest -Uri $asset.browser_download_url -Headers $headers -OutFile $archivePath
    Invoke-WebRequest -Uri $checksumAsset.browser_download_url -Headers $headers -OutFile $checksumPath

    $expectedHash = (Get-Content $checksumPath |
        Where-Object { $_ -match ("\s" + [regex]::Escape($asset.name) + "$") } |
        ForEach-Object { ($_ -split "\s+")[0] } |
        Select-Object -First 1)
    if ([string]::IsNullOrWhiteSpace($expectedHash)) {
        throw "No checksum was found for $($asset.name)."
    }

    $actualHash = (Get-FileHash -Algorithm SHA256 -Path $archivePath).Hash
    if ($actualHash -ne $expectedHash) {
        throw "Checksum verification failed for $($asset.name)."
    }

    Expand-Archive -Path $archivePath -DestinationPath $extractDir
    $binary = Join-Path $extractDir "kin.exe"
    if (-not (Test-Path $binary)) {
        throw "The release archive does not contain kin.exe."
    }

    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    Copy-Item -Force $binary (Join-Path $InstallDir "kin.exe")

    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    $pathEntries = @($userPath -split ";" | Where-Object { $_ })
    if ($pathEntries -notcontains $InstallDir) {
        [Environment]::SetEnvironmentVariable("Path", (($pathEntries + $InstallDir) -join ";"), "User")
    }

    Write-Host "Installed Kin to $(Join-Path $InstallDir 'kin.exe')."
    Write-Host "Open a new terminal, then run: kin init"
}
finally {
    if (Test-Path $tempDir) {
        Remove-Item -Recurse -Force $tempDir
    }
}
