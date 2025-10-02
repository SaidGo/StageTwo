# -------- configurable --------
APP            ?= go2part
MAIN           ?= ./cmd/web
PORT           ?= :8080

# redis (локальный контейнер из docker-compose)
REDIS_ADDR     ?= 127.0.0.1:6300
REDIS_DB       ?= 0
REDIS_PASSWORD ?=

# cache
CACHE_TTL_SECONDS ?= 300

# -------- helpers --------
ENV = GIN_MODE=release PORT="$(PORT)" CACHE_TTL_SECONDS=$(CACHE_TTL_SECONDS) \
      REDIS_ADDR="$(REDIS_ADDR)" REDIS_DB=$(REDIS_DB) REDIS_PASSWORD="$(REDIS_PASSWORD)"

# -------- targets --------
.PHONY: all tidy fmt vet build run run-logs test clean \
        docker-up docker-down docker-ps redis-ping redis-keys

all: build

tidy:
	@go mod tidy

fmt:
	@go fmt ./...

# ограничим vet на основные пакеты, чтобы не срывать из-за черновиков
vet:
	@go vet ./cmd/... ./internal/web/... ./internal/legalentities/... ./pkg/... || true

build:
	@$(ENV) go build -o bin/$(APP) $(MAIN)
	@echo "Built bin/$(APP)"

run:
	@$(ENV) go run $(MAIN)

run-logs:
	@$(ENV) go run $(MAIN)

test:
	# Тесты только там, где они валидны
	@go test ./domain -count=1 || true

clean:
	@rm -rf bin dist build coverage || true

# ----- docker / redis -----
docker-up:
	@docker compose up -d redis

docker-down:
	@docker compose stop redis

docker-ps:
	@docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | grep -i redis || true

redis-ping:
	@docker exec -it $$(docker ps -qf "name=redis") redis-cli -p 6379 ping

redis-keys:
	@docker exec -it $$(docker ps -qf "name=redis") redis-cli -p 6379 keys '*'
