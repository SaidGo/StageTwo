# Go2part

Модульный backend-проект на Go с архитектурой DDD, OpenAPI и интеграцией с Kafka.

## Стек

- **Go 1.22+**
- **Gin** — HTTP-сервер
- **GORM** — ORM
- **PostgreSQL** — основная БД
- **golang-migrate** — миграции
- **Kafka + Kafdrop** — события и их мониторинг

## Структура проекта

```
cmd/               # точки входа (web, cli, tools)
domain/            # доменные модели
dto/               # DTO для API
internal/          # бизнес-логика и сервисы (в т.ч. internal/kafka)
migrations/        # SQL-миграции
openapi/           # OpenAPI-спецификация
pkg/               # инфраструктурные пакеты
scripts/           # скрипты для запуска/отладки
docs/, grafana/    # доп. материалы и метрики
```

## Быстрый старт

Требования: **Go 1.22+**, **Docker + Docker Compose**, **make**, **bash**.

```bash
# 1) Поднять Kafka/ZooKeeper/Kafdrop
make up

# 2) Создать топики
make kafka-topics

# 3) Запустить веб-сервис (foreground)
make run-fore
# или в фоне
make run-bg
# посмотреть логи
make logs-web
```

## Переменные окружения

- `WEB_ADDR` — адрес HTTP-сервера (по умолчанию `:8080`)
- `KAFKA_BROKERS` — брокеры Kafka (по умолчанию `localhost:29092`)
- `POSTGRES_DSN` — строка подключения к PostgreSQL, например:
  ```
  postgres://postgres:password@localhost:5432/go2part?sslmode=disable
  ```

## API

Документация API описана в `openapi/openapi.yaml`.

Можно просмотреть через [Swagger Editor](https://editor.swagger.io/).

---

© Go2part Project
