$installDir = Join-Path $HOME ".min\bin"
$exeName = "min.exe"
$exePath = Join-Path $installDir $exeName
$url = "https://github.com/umastrodev12/Min/releases/latest/download/min.exe"

Write-Host "Starting Min CLI installation..."

if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Force -Path $installDir | Out-Null
}

Write-Host "Downloading min.exe..."
try {
    Invoke-WebRequest -Uri $url -OutFile $exePath -ErrorAction Stop
    Write-Host "[OK] Executable downloaded."
} catch {
    Write-Host "[ERROR] Failed to download min.exe. Make sure you have a Release with the file attached."
    exit
}

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notmatch [regex]::Escape($installDir)) {
    $newPath = "$userPath;$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "[OK] PATH variable updated!"
}

Write-Host "Installation completed successfully!"
Write-Host "Restart your terminal and type: min help"
