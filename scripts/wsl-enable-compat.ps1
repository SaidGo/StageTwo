$ErrorActionPreference = "Stop"
Get-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux | Format-Table FeatureName,State -Auto
Get-WindowsOptionalFeature -Online -FeatureName VirtualMachinePlatform | Format-Table FeatureName,State -Auto
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux -All -NoRestart | Out-Null
Enable-WindowsOptionalFeature -Online -FeatureName VirtualMachinePlatform -All -NoRestart | Out-Null
Write-Host "Если 'Enable Pending' — выполни Restart-Computer"
