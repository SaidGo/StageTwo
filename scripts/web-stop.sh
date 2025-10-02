#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PIDFILE="$ROOT/bin/web.pid"

# 1) стоп по PID-файлу (если есть)
if [ -f "$PIDFILE" ]; then
  PID="$(tr -d '\r' < "$PIDFILE" || true)"
  if [ -n "${PID:-}" ]; then
    /c/Windows/System32/taskkill.exe /PID "$PID" /F >/dev/null 2>&1 || true
  fi
  rm -f "$PIDFILE"
fi

# 2) добивка по имени
/c/Windows/System32/taskkill.exe /IM web.exe /F >/dev/null 2>&1 || true

# 3) очистка 8080
bash "$ROOT/scripts/port-kill-8080.sh" || true

echo "web stopped"
