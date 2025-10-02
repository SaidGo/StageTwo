#!/usr/bin/env bash
set -e
docker exec -it go2part-prometheus-1 /bin/sh -c '
  echo "[inside prometheus] resolv:"; getent hosts host.docker.internal || true
  echo "[inside prometheus] curl metrics:"; \
  (command -v wget >/dev/null 2>&1 && wget -S -O- http://host.docker.internal:8080/metrics 2>&1 | head -n 15) || echo "wget not available"
'
