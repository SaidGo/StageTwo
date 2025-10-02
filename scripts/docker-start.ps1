# # E:/Projects/Go2part/scripts/docker-start.ps1
$ErrorActionPreference = "Stop"

function Test-RebootPendingWSL {
  $need = $false
  $f1 = Get-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
  $f2 = Get-WindowsOptionalFeature -Online -FeatureName VirtualMachinePlatform
  if ($f1.State -eq 'Enable Pending' -or $f2.State -eq 'Enable Pending') { $need = $true }
  if ($need) { Write-Host "Нужна перезагрузка Windows (WSL/VirtualMachinePlatform = Enable Pending)."; exit 2 }
}

function Start-DockerServiceIfNeeded {
  try {
    $svc = Get-Service -Name "com.docker.service" -ErrorAction Stop
    if ($svc.Status -ne "Running") {
      Write-Host "Starting Windows service: com.docker.service ..."
      Start-Service -Name "com.docker.service"
      $svc.WaitForStatus("Running","00:00:30")
    }
  } catch { Write-Host "Service 'com.docker.service' отсутствует/нет прав. Запускаем GUI." }
}

function Start-DockerDesktopIfNeeded {
  $exe = "C:\Program Files\Docker\Docker\Docker Desktop.exe"
  if (-not (Get-Process -Name "Docker Desktop" -ErrorAction SilentlyContinue)) {
    if (-not (Test-Path $exe)) { throw "Docker Desktop не установлен: $exe" }
    Write-Host "Starting Docker Desktop GUI..."
    Start-Process -FilePath $exe | Out-Null
  }
}

function Wait-DockerEngine {
  Write-Host "Waiting for Docker engine..."
  for ($i=0; $i -lt 300; $i++) {
    try {
      $v = docker version --format '{{.Server.Version}}' 2>$null
      if ($LASTEXITCODE -eq 0 -and $v) { Write-Host "Docker engine is up. Version: $v"; return }
    } catch {}
    Start-Sleep -Seconds 2
  }
  Write-Host "Docker engine did not become ready in time."
  exit 1
}

Test-RebootPendingWSL
Start-DockerServiceIfNeeded
Start-DockerDesktopIfNeeded
Wait-DockerEngine

Write-Host "`ndocker version:"; docker version
Write-Host "`ndocker compose version:"; docker compose version

