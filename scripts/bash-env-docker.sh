#!/usr/bin/env bash
set +e
export PATH="/c/Program Files/Docker/Docker/resources/bin:$PATH"
export PATH="/c/Program Files/Docker/Docker/cli-plugins:$PATH"
if ! command -v docker >/dev/null 2>&1; then
  echo "[warn] docker.exe не найден в PATH"
else
  docker version >/dev/null 2>&1 || echo "[warn] docker CLI ок, daemon ещё поднимается"
fi
