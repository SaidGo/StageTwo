#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

export REDIS_ADDR=${REDIS_ADDR:-127.0.0.1:6300}
export CACHE_TTL_SECONDS=${CACHE_TTL_SECONDS:-300}
export CACHE_LOG_MISS=0

echo "web starting; REDIS_ADDR=$REDIS_ADDR TTL=${CACHE_TTL_SECONDS}s CACHE_LOG_MISS=$CACHE_LOG_MISS"
cd "$ROOT"
./bin/web > "$ROOT/bin/web.log" 2>&1 &
echo "web started; pid=$!; logs: $ROOT/bin/web.log"
