# Go2part — Legal Entity CRUD (Этап 2.1)

## 📌 Цель
Реализация CRUD для юридических лиц (`Legal Entity`) в рамках архитектуры:

Federation
└── Company
└── Legal Entity
└── Bank Account (в будущем)

markdown
Копировать

## 🏗️ Структура проекта

- `internal/legalentities/`
  - `orm.go` — описание модели
  - `repository.go` — доступ к БД
  - `main.go` — сервис
- `internal/web/legal_entity_handler.go` — HTTP-хендлеры
- `internal/router/routers-legal-entities.go` — маршруты
- `migrations/000001_create_legal_entities.*.sql` — SQLite миграции
- `internal/web/olegalentity/` — автогенерируемый API-интерфейс по OpenAPI
- `openapi/openapi.yaml` — описание API

## 🧪 Примеры запросов (Postman)

### 🔹 Создание
```http
POST /legal-entities
Content-Type: application/json

{
  "name": "ООО Пример"
}
🔹 Получение всех
http
Копировать
GET /legal-entities
🔹 Обновление
http
Копировать
PUT /legal-entities/{uuid}
Content-Type: application/json

{
  "name": "АО Обновлено"
}
🔹 Удаление
http
Копировать
DELETE /legal-entities/{uuid}
⚙️ Сборка и запуск
bash
Копировать
# Генерация wire
go generate ./...

# Сборка
go build -o web.exe ./cmd/web

# Запуск
./web.exe
🗃️ Миграции
bash
Копировать
# Применить миграции
./migrate.exe -database "sqlite3://E:/Projects/Go2part/legalentities.db" -path ./migrations up

# Проверить таблицу
sqlite3 legalentities.db ".tables"
🧱 Зависимости
Go >= 1.21

SQLite

migrate (CLI)

🧾 Лицензия
MIT