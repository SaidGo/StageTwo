.PHONY: clean-bak clean fmt build run migrate-up migrate-down gen swagger

APP_BIN := ./bin/web
PKG_DIRS := ./internal ./dto ./domain

# -------- Formatting --------
fmt:
	gofmt -w $(PKG_DIRS)

# -------- Build/Run --------
build:
	go mod tidy
	mkdir -p ./bin
	go build -o $(APP_BIN) ./cmd/web

run:
	@if [ -z "$$POSTGRES_DSN" ]; then \
		echo "POSTGRES_DSN is not set"; \
		exit 1; \
	fi
	$(APP_BIN)

# -------- Migrations (golang-migrate совместимый бинарь) --------
# Переменные среды:
# MIGRATE_BIN=./migrate_bin/migrate
# POSTGRES_DSN=postgres://postgres:pass@localhost:5432/go2part?sslmode=disable
migrate-up:
	@if [ -z "$$MIGRATE_BIN" ]; then echo "Set MIGRATE_BIN"; exit 1; fi
	@if [ -z "$$POSTGRES_DSN" ]; then echo "Set POSTGRES_DSN"; exit 1; fi
	"$$MIGRATE_BIN" -path "./migrations" -database "$$POSTGRES_DSN" up

migrate-down:
	@if [ -z "$$MIGRATE_BIN" ]; then echo "Set MIGRATE_BIN"; exit 1; fi
	@if [ -z "$$POSTGRES_DSN" ]; then echo "Set POSTGRES_DSN"; exit 1; fi
	"$$MIGRATE_BIN" -path "./migrations" -database "$$POSTGRES_DSN" down 1

# -------- OpenAPI generation (если используете генератор) --------
gen:
	@echo "Place your openapi generation commands here"
	@echo "e.g. oapi-codegen -generate types,server -o internal/web/generated.go openapi/openapi.yaml"

swagger:
	@echo "Open openapi/openapi.yaml in Swagger Editor or run local UI as needed"

# -------- Cleanup --------
clean-bak:
	@find . -type f \( -name "*.bak" -o -name "*.bak_old" -o -name "*.tmp" \) -print -delete

clean:
	rm -rf ./bin
