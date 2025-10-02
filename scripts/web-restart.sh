#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
bash "$ROOT/scripts/web-stop.sh" || true
"$ROOT/bin/web" >"$ROOT/bin/web.log" 2>&1 &
echo $! >"$ROOT/bin/web.pid"
echo "web started; pid=$(cat "$ROOT/bin/web.pid"); logs: $ROOT/bin/web.log"
