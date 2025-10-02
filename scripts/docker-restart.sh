#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

"$ROOT/scripts/win-task.sh" taskkill /IM com.docker.backend.exe /F || true
"$ROOT/scripts/win-task.sh" taskkill /IM "Docker Desktop.exe" /F || true

DOCKER_DESKTOP="C:\\Program Files\\Docker\\Docker\\Docker Desktop.exe"
"$ROOT/scripts/win-task.sh" start "$DOCKER_DESKTOP"

echo "waiting for Docker daemon ..."
for i in $(seq 1 60); do
  if docker info > /dev/null 2>&1; then
    echo "Docker is ready"
    exit 0
  fi
  sleep 2
done

echo "Docker daemon not ready (timeout)" >&2
exit 1
