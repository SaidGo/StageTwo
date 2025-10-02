#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"

BASE="${BASE:-http://127.0.0.1:8080}"
LE_PATH="${LE_PATH:-/legal-entities}"

# два запроса: MISS -> HIT
curl -fsS -o /dev/null "${BASE}${LE_PATH}" || true
curl -fsS -o /dev/null "${BASE}${LE_PATH}" || true

# прямые проверки в redis
echo "TTL cache:legal_entities:"
"${CLI}" TTL "cache:legal_entities" | tr -d '\r' || true

echo "EXISTS cache:legal_entities:"
"${CLI}" EXISTS "cache:legal_entities" | tr -d '\r' || true

echo "KEYS cache:* (debug):"
"${CLI}" KEYS 'cache:*' || true
