
Write-Host "Downloading drs..."

$url = "https://github.com/arghyagod-coder/GoLang/releases/latest/download/drs-windows-amd64.exe"

$dir = $env:USERPROFILE + "\.drs"
$filepath = $env:USERPROFILE + "\.drs\drs.exe"

[System.IO.Directory]::CreateDirectory($dir)
(Invoke-WebRequest -Uri $url -OutFile $filepath)

Write-Host "Adding drs to PATH..."
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";"+$dir,
    [EnvironmentVariableTarget]::Machine)

Write-Host 'drs installation is successfull!'
Write-Host "You need to restart your shell to use drs."
