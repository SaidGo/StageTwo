# # E:/Projects/Go2part/scripts/start-monitoring.ps1
$ErrorActionPreference = "Stop"

# 1) �������� Docker CLI � ������
try { docker version | Out-Null } catch { throw "docker daemon �� �������" }

# 2) ����� ������� � compose
Set-Location "E:\Projects\Go2part"

# 3) ������� ����������
docker compose up -d prometheus grafana

# 4) ���������
docker compose ps

Write-Host ""
Write-Host "Prometheus: http://localhost:9090/targets"
Write-Host "Grafana:    http://localhost:3000  (admin/admin)"

