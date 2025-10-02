#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"
KEYS="${ROOT}/scripts/cache-keys.sh"
LOG="${ROOT}/bin/web.log"

# 0) ENV (опционально)
: "${CACHE_LOG_MISS:=0}"
export CACHE_LOG_MISS

# 1) PURGE
"${ROOT}/scripts/cache-purge.sh" || true

# 2) Прогрев
curl -s "http://127.0.0.1:8080/legal-entities"  > /dev/null
curl -s "http://127.0.0.1:8080/bank_accounts"   > /dev/null

# 3) Повторные запросы (должны дать HIT)
curl -s "http://127.0.0.1:8080/legal-entities"  > /dev/null
curl -s "http://127.0.0.1:8080/bank_accounts"   > /dev/null

# 4) Ключи и TTL
"${KEYS}"

# 5) Логи HIT
echo "---- cache HIT lines (last 200) ----"
tail -n 200 "${LOG}" | grep -E "cache: (SET|HIT|DEL) cache:(legal_entities|bank_accounts)" || true
