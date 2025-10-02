$ErrorActionPreference = "SilentlyContinue"
$exe = "C:\Program Files\Docker\Docker\Docker Desktop.exe"
# Запустить GUI, если не запущен
if (-not (Get-Process -Name "Docker Desktop" -ErrorAction SilentlyContinue)) {
  Start-Process -FilePath $exe | Out-Null
}
# Ждём доступности демона
Write-Host "Waiting for Docker engine..."
for ($i=0; $i -lt 300; $i++) {
  try {
    $v = docker version --format '{{.Server.Version}}' 2>$null
    if ($LASTEXITCODE -eq 0 -and $v) { Write-Host "Docker engine is up. Version: $v"; exit 0 }
  } catch {}
  Start-Sleep -Seconds 2
}
Write-Error "Docker engine did not become ready in time."
