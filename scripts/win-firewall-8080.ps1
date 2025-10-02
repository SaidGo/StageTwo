$ErrorActionPreference = "Stop"
$rule = Get-NetFirewallRule -DisplayName "Go2part 8080 inbound" -ErrorAction SilentlyContinue
if (-not $rule) {
  New-NetFirewallRule -DisplayName "Go2part 8080 inbound" `
    -Direction Inbound -Action Allow -Protocol TCP -LocalPort 8080 `
    -Profile Any | Out-Null
  Write-Host "Firewall rule created"
} else {
  Write-Host "Firewall rule already exists"
}
