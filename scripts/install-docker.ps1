# Требования: PowerShell от имени Администратора

# 1) Проверка прав
$principal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
if (-not $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
  Write-Error "Запусти PowerShell от имени Администратора."
  exit 1
}

# 2) Включение компонентов WSL2
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux -NoRestart | Out-Null
Enable-WindowsOptionalFeature -Online -FeatureName VirtualMachinePlatform -NoRestart | Out-Null

# 3) Выбор WSL2 по умолчанию
wsl --set-default-version 2

# 4) Установка Docker Desktop (через winget)
# Если winget отсутствует — обнови App Installer из Microsoft Store.
winget install -e --id Docker.DockerDesktop --accept-source-agreements --accept-package-agreements

Write-Host "`nУстановка завершена. Перезагрузи Windows, запусти Docker Desktop и дождись статуса 'Running'."
Write-Host "После запуска проверь в новом окне терминала: docker version && docker compose version"

