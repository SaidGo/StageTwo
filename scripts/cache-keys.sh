#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"

echo "Keys:"
"${CLI}" KEYS 'cache:*' | tr -d '\r' | while read -r k; do
  [[ -z "$k" ]] && continue
  ttl="$("${CLI}" TTL "$k" | tr -d '\r')"
  echo "  $k  ttl=${ttl}s"
done
