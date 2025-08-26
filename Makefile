
APP_NAME=web
BINARY=bin/$(APP_NAME)
PKG=example.com/local/Go2part

.PHONY: all build run clean wire migrate fmt lint

all: build

## Build application
build:
	go build -o $(BINARY) ./cmd/web

## Run application
run:
	go run ./cmd/web

## Wire dependency injection
wire:
	cd internal/app && wire

## Run migrations (Postgres DSN must be set in POSTGRES_DSN)
migrate:
	migrate -path migrations -database "$$POSTGRES_DSN" up

## Format code
fmt:
	gofmt -s -w .

## Lint (staticcheck required)
lint:
	staticcheck ./...

## Clean build artifacts
clean:
	rm -rf bin/ web.exe *.out *.test
