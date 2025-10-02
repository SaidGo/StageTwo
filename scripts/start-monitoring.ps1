# # E:/Projects/Go2part/scripts/start-monitoring.ps1
$ErrorActionPreference = "Stop"

# 1) Проверка Docker CLI и демона
try { docker version | Out-Null } catch { throw "docker daemon не запущен" }

# 2) Папка проекта и compose
Set-Location "E:\Projects\Go2part"

# 3) Поднять мониторинг
docker compose up -d prometheus grafana

# 4) Состояние
docker compose ps

Write-Host ""
Write-Host "Prometheus: http://localhost:9090/targets"
Write-Host "Grafana:    http://localhost:3000  (admin/admin)"

