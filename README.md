# Go2part

## Описание
Go2part — backend-сервис, реализующий многоуровневую архитектуру **Federation → Company → Legal Entity → Bank Account** с использованием Go, GORM, Gin, OpenAPI и миграций.

## Структура проекта
```
cmd/               # CLI и Web точки входа
domain/            # Доменные сущности
dto/               # DTO для API
internal/          # Логика по доменам
  ├─ legalentities # CRUD Legal Entity
  ├─ company       # Логика компаний
  ├─ federation    # Логика федераций
  ├─ web           # OpenAPI-хендлеры и middleware
migrations/        # SQL-миграции
openapi/           # OpenAPI-описания
pkg/               # Инфраструктурные пакеты (postgres, cache, redis и др.)
scripts/           # Скрипты генерации и обслуживания
```

## Основные команды
```bash
# Сборка web-сервера
go build ./cmd/web

# Сборка CLI
go build ./cmd/cli

# Генерация OpenAPI
make gen-openapi

# Применение миграций (SQLite)
make migrate-up-sqlite

# Применение миграций (PostgreSQL)
make migrate-up-pg

# Очистка
make clean

# Генерация Wire
make wire
```

## Пример запуска (SQLite)
```bash
make build
make migrate-up-sqlite
./web
```

## Требования
- Go >= 1.21
- SQLite3 / PostgreSQL
- make / bash-окружение (MSYS под Windows)
- Wire
- OpenAPI Generator
