#!/usr/bin/env bash
set -euo pipefail

BASE="${BASE:-http://127.0.0.1:8080}"
LE_PATH="${LE_PATH:-/legal-entities}"
BA_PATH="${BA_PATH:-/bank_accounts}"

try200() {
  local url="$1" name="$2" n=0
  while true; do
    code="$(curl -fsS -o /dev/null -w "%{http_code}" "$url" || echo 000)"
    if [[ "$code" == "200" || "$code" == "204" ]]; then
      break
    fi
    n=$((n+1))
    if [[ $n -gt 60 ]]; then
      echo "FAIL: $name not ready (last HTTP $code): $url" >&2
      exit 1
    fi
    sleep 0.5
  done
}

# 1) дождаться готовности по реальным маршрутам
try200 "${BASE}${LE_PATH}" "legal-entities"
try200 "${BASE}${BA_PATH}" "bank_accounts"

# 2) прогрев (двойные вызовы -> HIT после MISS)
curl -fsS -o /dev/null "${BASE}${LE_PATH}" || true
curl -fsS -o /dev/null "${BASE}${LE_PATH}" || true

curl -fsS -o /dev/null "${BASE}${BA_PATH}" || true
curl -fsS -o /dev/null "${BASE}${BA_PATH}" || true

echo "warm done"
