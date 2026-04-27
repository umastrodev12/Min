@echo off
set "targetDir=%USERPROFILE%\.min\bin"
if not exist "%targetDir%" mkdir "%targetDir%"

copy "min.exe" "%targetDir%\min.exe"

setx PATH "%PATH%;%targetDir%"

echo Min CLI installed successfully! 
echo Please restart your terminal and type 'min help'.
pause