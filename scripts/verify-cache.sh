#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
KEYS="${ROOT}/scripts/cache-keys.sh"
WARM="${ROOT}/scripts/warm.sh"
CLI="${ROOT}/scripts/redis-cli.sh"
BASE="${BASE:-http://127.0.0.1:8080}"

# 1) Прогрев (ждёт готовность реальных маршрутов)
"${WARM}" >/dev/null

# 2) Проверка наличия ключей
echo "[verify] Redis keys:"
"${KEYS}"

le_exists="$("${CLI}" EXISTS "cache:legal_entities" | tr -d '\r')"
ba_exists="$("${CLI}" EXISTS "cache:bank_accounts" | tr -d '\r')"
le_ttl="$("${CLI}" TTL "cache:legal_entities" | tr -d '\r' || echo -1)"
ba_ttl="$("${CLI}" TTL "cache:bank_accounts" | tr -d '\r' || echo -1)"

echo "[verify] cache:legal_entities EXISTS=${le_exists} TTL=${le_ttl}"
echo "[verify] cache:bank_accounts  EXISTS=${ba_exists} TTL=${ba_ttl}"

# 3) Критерии прохождения
fail=0
[[ "${le_exists}" != "1" || "${le_ttl}" -le 0 ]] && fail=1
[[ "${ba_exists}" != "1" || "${ba_ttl}" -le 0 ]] && fail=1

# 4) Логи (HIT для обоих ключей)
echo "[verify] recent HIT logs:"
grep -a -E "cache: HIT cache:(legal_entities|bank_accounts)" "${ROOT}/bin/web.log" | tail -n 10 || true

if [[ $fail -ne 0 ]]; then
  echo "[verify] FAIL" >&2
  exit 1
fi
echo "[verify] OK"
