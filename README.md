# Go2part

Модульный backend-проект на Go с архитектурой DDD и поддержкой OpenAPI-генерации.

## Стек

- **Go 1.22+**
- **Gin** (http-сервер)
- **GORM** (ORM)
- **Postgres** (основная база данных)
- **golang-migrate** (миграции)

## Структура проекта

```
cmd/               # точки входа (web, cli)
domain/            # доменные модели
dto/               # DTO для API
internal/          # бизнес-логика и сервисы
migrations/        # SQL-миграции
openapi/           # OpenAPI-спецификация
pkg/               # инфраструктурные пакеты
scripts/           # bash-скрипты для сборки и тестов
```

## Запуск

### Требования

- Go 1.22+
- PostgreSQL 14+
- make, bash

### Команды

```bash
# Сборка
make build

# Запуск веб-сервера
make run

# Запуск миграций
make migrate-up

# Откат миграций
make migrate-down

# Генерация OpenAPI-клиентов и серверов
make gen

# Очистка
make clean
```

## Переменные окружения

- `POSTGRES_DSN` — строка подключения к PostgreSQL, например:
  ```
  postgres://postgres:password@localhost:5432/go2part?sslmode=disable
  ```

## API

Документация API описана в `openapi/openapi.yaml`.

Можно просмотреть через [Swagger Editor](https://editor.swagger.io/).

---

© Go2part Project
