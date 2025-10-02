#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PIDFILE="$ROOT/bin/web.pid"

# kill by pidfile if exists
if [ -f "$PIDFILE" ]; then
  PID="$(cat "$PIDFILE" || true)"
  if [ -n "${PID:-}" ]; then
    /c/Windows/System32/taskkill.exe /PID "$PID" /F >/dev/null 2>&1 || true
  fi
  rm -f "$PIDFILE"
fi

# kill anything still holding 8080
/c/Windows/System32/netstat.exe -ano | grep ":8080" | awk '{print $5}' | sed -E 's/.*://g' | xargs -r -I{} /c/Windows/System32/taskkill.exe /PID {} /F >/dev/null 2>&1 || true
echo "port 8080 cleared"
