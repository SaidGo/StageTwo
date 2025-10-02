#!/usr/bin/env bash
set -euo pipefail

cmd() {
  /c/Windows/System32/cmd.exe /C "$*"
}

case "${1-}" in
  tasklist)
    shift
    cmd tasklist "$@"
    ;;
  taskkill)
    shift
    cmd taskkill "$@"
    ;;
  start)
    shift
    cmd start "" "$@"
    ;;
  *)
    echo "usage: $0 {tasklist|taskkill|start} <args...>" >&2
    exit 2
    ;;
esac
