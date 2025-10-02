#!/usr/bin/env bash
set -euo pipefail

CID="${1:-go2part-kafka-1}"

docker exec -it "$CID" bash -lc '
set -e
kafka-topics --bootstrap-server localhost:9092 --create --if-not-exists --topic legal-entities-created --replication-factor 1 --partitions 1
kafka-topics --bootstrap-server localhost:9092 --create --if-not-exists --topic bank-accounts-created --replication-factor 1 --partitions 1
echo "== topics =="
kafka-topics --bootstrap-server localhost:9092 --list | sort
'
