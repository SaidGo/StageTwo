#!/usr/bin/env bash
set -euo pipefail
command -v docker >/dev/null 2>&1 || { echo "docker не найден в PATH"; exit 1; }
docker version >/dev/null 2>&1 || { echo "docker daemon не запущен"; exit 1; }
cd /e/Projects/Go2part
docker compose up -d prometheus grafana
docker compose ps
echo
echo "Prometheus: http://localhost:9090/targets"
echo "Grafana:    http://localhost:3000  (admin/admin)"
