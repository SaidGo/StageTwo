#!/usr/bin/env bash
set +e

echo "== Warmup =="
curl -s -o /dev/null -w "GET /legal-entities -> %{http_code}\n"  http://localhost:8080/legal-entities
curl -s -o /dev/null -w "GET /bank_accounts  -> %{http_code}\n"  http://localhost:8080/bank_accounts

echo "== App counters (raw /metrics) =="
curl -s http://localhost:8080/metrics | grep -E 'legal_entities_requests_total|bank_accounts_requests_total' || echo "[warn] counters not found"

echo "== Wait first scrape (>=5s) =="
sleep 6

echo "== Prometheus API: active targets (short) =="
curl -s 'http://localhost:9090/api/v1/targets?state=active' \
| tr -d '\r' \
| awk '
  /"activeTargets":\[/,/\]}/ { buf=buf $0 }
  END {
    n = split(buf, arr, /{/)
    for (i=1;i<=n;i++) if (arr[i] ~ /"job":"app"/) {
      match(arr[i], /"instance":"[^"]+"/, inst)
      match(arr[i], /"health":"[^"]+"/, health)
      match(arr[i], /"lastError":"[^"]*"/, err)
      if (inst[0]!="") print inst[0], health[0], err[0]
    }
  }'

echo "== Prometheus API: up{job=\"app\"} =="
curl -sG --data-urlencode 'query=up{job="app"}' http://localhost:9090/api/v1/query
echo

echo "== Prometheus API: totals =="
curl -sG --data-urlencode 'query=sum(legal_entities_requests_total)'  http://localhost:9090/api/v1/query; echo
curl -sG --data-urlencode 'query=sum(bank_accounts_requests_total)'   http://localhost:9090/api/v1/query; echo

echo "== Prometheus API: RPS (rate[1m]) =="
curl -sG --data-urlencode 'query=sum(rate(legal_entities_requests_total[1m]))'  http://localhost:9090/api/v1/query; echo
curl -sG --data-urlencode 'query=sum(rate(bank_accounts_requests_total[1m]))'   http://localhost:9090/api/v1/query; echo
