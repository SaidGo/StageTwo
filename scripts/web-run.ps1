# # E:/Projects/Go2part/scripts/web-run.ps1
$ErrorActionPreference = "Stop"

$proj = "E:\Projects\Go2part"
$bin  = Join-Path $proj "bin\web.exe"  # msys/go ������ web.exe, ��������� ����� ��� bin\web
$url  = "http://localhost:8080/metrics"

function Stop-Port([int]$Port) {
  try {
    $conns = Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue
    if ($conns) {
      $pids = ($conns | Select-Object -ExpandProperty OwningProcess) | Sort-Object -Unique
      foreach ($pid in $pids) { Stop-Process -Id $pid -Force -ErrorAction SilentlyContinue }
    }
  } catch { }
}

# 0) ����� ������ ������� �� :8080
Stop-Port 8080

# 1) ������ � ������ ���� ���� go � PATH
$go = Get-Command go -ErrorAction SilentlyContinue
if ($go) {
  Set-Location $proj
  & $go.Path build -o .\bin\web .\cmd\web
} else {
  if (-not (Test-Path $bin)) { throw "���������� ��� Go: ��� $bin. ������ �� Git Bash: go build -o ./bin/web ./cmd/web" }
}

# 2) SQLite-������
Remove-Item Env:POSTGRES_DSN -ErrorAction SilentlyContinue

# 3) ����� ������
$proc = Start-Process -FilePath $bin -WorkingDirectory $proj -PassThru
Write-Host ("web started, pid={0}" -f $proc.Id)

# 4) �������� /metrics
for ($i=0; $i -lt 60; $i++) {
  try {
    $r = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 2
    if ($r.StatusCode -eq 200) { Write-Host "/metrics OK (200)"; break }
  } catch { }
  Start-Sleep -Seconds 1
}

