#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"

LE="${1:-}"
if [[ -z "$LE" ]]; then
  echo "Usage: $0 <LE-UUID>" >&2
  exit 2
fi

BASE="${BASE:-http://127.0.0.1:8080}"
PATH_LE_BA="${PATH_LE_BA:-/legal-entities/${LE}/bank-accounts}"

curl -fsS -o /dev/null "${BASE}${PATH_LE_BA}" || true
curl -fsS -o /dev/null "${BASE}${PATH_LE_BA}" || true

echo "TTL cache:bank_accounts:${LE}:"
"${CLI}" TTL "cache:bank_accounts:${LE}" | tr -d '\r' || true

echo "EXISTS cache:bank_accounts:${LE}:"
"${CLI}" EXISTS "cache:bank_accounts:${LE}" | tr -d '\r' || true
