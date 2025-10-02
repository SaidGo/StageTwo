#!/usr/bin/env bash
set -euo pipefail
echo "[wait] docker daemon..."
for i in {1..60}; do
  if docker info >/dev/null 2>&1; then
    echo "[ok] docker ready"
    exit 0
  fi
  sleep 2
done
echo "[fail] docker not ready"; exit 1
