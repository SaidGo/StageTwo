#!/usr/bin/env bash
set -euo pipefail

CID="$(docker ps --format '{{.Names}}' | awk '/redis/ {print $1; exit}')"
if [[ -z "${CID}" ]]; then
  echo "Redis container not found. Start it: docker-compose -f \"E:/Projects/Go2part/docker-compose.yaml\" up -d redis" >&2
  exit 1
fi

REDIS_PASSWORD="${REDIS_PASSWORD-}"

if [[ -n "${REDIS_PASSWORD}" ]]; then
  exec docker exec -i "${CID}" redis-cli -a "${REDIS_PASSWORD}" "$@"
else
  exec docker exec -i "${CID}" redis-cli "$@"
fi
