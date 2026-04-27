$installDir = Join-Path $HOME ".min\bin"
$exeName = "min.exe"
$sourceExe = Join-Path $PSScriptRoot $exeName

Write-Host "Starting Min CLI installation..."

if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Force -Path $installDir | Out-Null
}

if (Test-Path $sourceExe) {
    Copy-Item -Path $sourceExe -Destination (Join-Path $installDir $exeName) -Force
    Write-Host "[OK] Executable copied."
} else {
    Write-Host "[ERROR] min.exe not found! Build the project first."
    Pause
    exit
}

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notmatch [regex]::Escape($installDir)) {
    $newPath = "$userPath;$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "[OK] PATH variable updated!"
} else {
    Write-Host "[OK] Min is already in your PATH."
}

Write-Host "Installation completed successfully!"
Write-Host "Close this terminal, open a new one and type: min version"
Write-Host "Press any key to exit..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")