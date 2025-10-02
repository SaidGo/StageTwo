#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"

mapfile -t KEYS < <("${CLI}" KEYS 'cache:*' | tr -d '\r')
if [[ ${#KEYS[@]} -eq 0 ]]; then
  echo "No cache:* keys."
  exit 0
fi

echo "Deleting ${#KEYS[@]} keys:"
for k in "${KEYS[@]}"; do
  echo "  DEL ${k}"
  "${CLI}" DEL "${k}" >/dev/null
done
echo "Done."
