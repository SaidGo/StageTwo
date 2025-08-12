.PHONY: build run test tidy wire migrate-up

BIN_DIR := bin

build:
	go build -o $(BIN_DIR)/web cmd/web/main.go

run:
	go run cmd/web/main.go

test:
	go test -v ./...

tidy:
	go mod tidy

wire:
	wire ./internal/app

# Требуется экспортировать POSTGRES_DSN с паролем:
#   export POSTGRES_DSN='postgres://postgres:Salavdi1@127.0.0.1:5432/go2part?sslmode=disable'
migrate-up:
	@if [ -z "$$POSTGRES_DSN" ]; then \
	  echo "Нужно экспортировать POSTGRES_DSN с паролем (postgres://user:pass@host:port/db?sslmode=disable)"; \
	  exit 1; \
	fi
	./migrate_bin/migrate.exe -path "./migrations" -database "$(POSTGRES_DSN)" up
