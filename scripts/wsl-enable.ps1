$ErrorActionPreference = "Stop"
$features = @('Microsoft-Windows-Subsystem-Linux','VirtualMachinePlatform')
Get-WindowsOptionalFeature -Online -FeatureName $features | Format-Table FeatureName,State -Auto
Enable-WindowsOptionalFeature -Online -FeatureName 'Microsoft-Windows-Subsystem-Linux' -All -NoRestart | Out-Null
Enable-WindowsOptionalFeature -Online -FeatureName 'VirtualMachinePlatform' -All -NoRestart | Out-Null
Write-Host "Если состояние = 'Enable Pending' — выполни: Restart-Computer"
