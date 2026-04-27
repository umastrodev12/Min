$installDir = Join-Path $HOME ".min\bin"
$exeName = "min.exe"
$exePath = Join-Path $installDir $exeName
$sourceExe = Join-Path $PSScriptRoot $exeName
$url = "https://github.com/umastrodev12/Min/releases/latest/download/min.exe"

Write-Host "Starting Min CLI installation..."

if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Force -Path $installDir | Out-Null
}

if (Test-Path $sourceExe) {
    Write-Host "[OK] Found min.exe locally. Copying..."
    Copy-Item -Path $sourceExe -Destination $exePath -Force
} else {
    Write-Host "[INFO] min.exe not found locally. Downloading from GitHub..."
    try {
        Invoke-WebRequest -Uri $url -OutFile $exePath -ErrorAction Stop
        Write-Host "[OK] Download successful."
    } catch {
        Write-Host "[ERROR] Could not download min.exe."
        Pause
        exit
    }
}

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notmatch [regex]::Escape($installDir)) {
    $newPath = "$userPath;$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "[OK] PATH variable updated!"
}

Write-Host "Installation completed successfully!"
Write-Host "Restart your terminal and type: min version"
Write-Host "Press any key to exit..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
