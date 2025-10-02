#!/usr/bin/env bash
set +e
echo "== Warmup =="
curl -s -o /dev/null -w "GET /legal-entities -> %{http_code}\n"  http://localhost:8080/legal-entities
curl -s -o /dev/null -w "GET /bank_accounts  -> %{http_code}\n"  http://localhost:8080/bank_accounts
echo "== App counters =="
curl -s http://localhost:8080/metrics | grep -E 'legal_entities_requests_total|bank_accounts_requests_total' || echo "[warn] counters not found"
echo "== Prometheus targets =="
curl -s http://localhost:9090/targets | grep -E 'job.*app' -m1 || echo "[warn] targets page not reachable"
echo "== Prometheus queries =="
curl -s "http://localhost:9090/api/v1/query?query=legal_entities_requests_total" | cut -c1-160; echo
curl -s "http://localhost:9090/api/v1/query?query=bank_accounts_requests_total"  | cut -c1-160; echo
curl -s "http://localhost:9090/api/v1/query?query=rate(legal_entities_requests_total[1m])" | cut -c1-160; echo
curl -s "http://localhost:9090/api/v1/query?query=rate(bank_accounts_requests_total[1m])"  | cut -c1-160; echo
