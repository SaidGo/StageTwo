$ErrorActionPreference = "SilentlyContinue"
Write-Host "[docker] stopping..."
Stop-Process -Name "com.docker.backend" -Force
Stop-Process -Name "Docker Desktop" -Force
Start-Sleep -Seconds 2
Write-Host "[docker] starting..."
Start-Process -FilePath "$Env:ProgramFiles\Docker\Docker\Docker Desktop.exe"
Write-Host "[docker] waiting..."
$retries = 60
for ($i=0; $i -lt $retries; $i++) {
  $p = (Get-Process -Name "com.docker.backend" -ErrorAction SilentlyContinue)
  if ($p) {
    $ver = (docker version 2>$null)
    if ($LASTEXITCODE -eq 0) { Write-Host "[docker] ready"; exit 0 }
  }
  Start-Sleep -Seconds 2
}
Write-Host "[docker] not ready"; exit 1
