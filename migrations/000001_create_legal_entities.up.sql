-- SQLite / универсальная схема
CREATE TABLE IF NOT EXISTS legal_entities (
  uuid TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_legal_entities_deleted_at ON legal_entities(deleted_at);

-- Для PostgreSQL используйте:
-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";
-- CREATE TABLE IF NOT EXISTS legal_entities (
--   uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--   name TEXT NOT NULL,
--   created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
--   updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
--   deleted_at TIMESTAMPTZ NULL
-- );
-- CREATE INDEX IF NOT EXISTS idx_legal_entities_deleted_at ON legal_entities(deleted_at);
