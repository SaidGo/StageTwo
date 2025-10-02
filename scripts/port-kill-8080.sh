#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

PIDS=$(/c/Windows/System32/netstat.exe -ano \
  | tr -d '\r' \
  | awk '$2 ~ /:8080$/ && $6 == "LISTENING" {print $5} $2 ~ /\[::\]:8080$/ && $6 == "LISTENING" {print $5}' \
  | sort -u)

if [ -z "${PIDS:-}" ]; then
  echo "8080: свободен"
  exit 0
fi

echo "kill PIDs: ${PIDS}"
for pid in ${PIDS}; do
  "$ROOT/scripts/win-task.sh" taskkill /PID "${pid}" /F || true
done

LEFT=$(/c/Windows/System32/netstat.exe -ano | tr -d '\r' | awk '$2 ~ /:8080$/ && $6 == "LISTENING" || $2 ~ /\[::\]:8080$/ && $6 == "LISTENING"')
if [ -n "${LEFT}" ]; then
  echo "8080 всё ещё занят:"
  echo "${LEFT}"
  exit 1
fi
echo "8080 освобождён"
