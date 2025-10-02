#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI="${ROOT}/scripts/redis-cli.sh"

"${CLI}" PING
"${CLI}" KEYS 'cache:*' || true

if [[ "${1-}" != "" ]]; then
  "${CLI}" TTL "$1" || true
fi
