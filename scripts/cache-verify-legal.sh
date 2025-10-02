#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"
KEYS="${ROOT}/scripts/cache-keys.sh"

# прогрев
curl -s "http://127.0.0.1:8080/legal-entities"  > /dev/null
curl -s "http://127.0.0.1:8080/legal-entities"  > /dev/null

# проверка ключа и TTL
"${KEYS}"

# прямой TTL по ключу
"${CLI}" TTL "cache:legal_entities" | tr -d '\r'
