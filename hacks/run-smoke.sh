#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.."; pwd)"
cd "$ROOT"

echo "[1/4] build"
go build -o bin/web ./cmd/web

echo "[2/4] migrate (SQLite)"
rm -f legalentities.db
sqlite3 legalentities.db < migrations/000001_create_legal_entities.up.sql

echo "[3/4] start server"
./bin/web & SERVER_PID=$!
trap 'kill $SERVER_PID 2>/dev/null || true' EXIT
sleep 1

echo "[4/4] smoke CRUD"
POST=$(curl -s -X POST 'http://127.0.0.1:8080/legal-entities' \
  -H 'Content-Type: application/json' \
  -d '{"name":"Smoke Co"}')
echo "POST: $POST"

ID=$(echo "$POST" | sed -n 's/.*"uuid":"\([^"]*\)".*/\1/p')
echo "UUID: $ID"

echo "LIST:"; curl -s 'http://127.0.0.1:8080/legal-entities?limit=5&offset=0'; echo
echo "GET :" ; curl -s "http://127.0.0.1:8080/legal-entities/$ID"; echo
echo "PUT :" ; curl -s -X PUT "http://127.0.0.1:8080/legal-entities/$ID" \
  -H 'Content-Type: application/json' -d '{"name":"Smoke Co Updated"}'; echo
echo "DEL :" ; curl -s -o /dev/null -w '%{http_code}\n' -X DELETE \
  "http://127.0.0.1:8080/legal-entities/$ID"
