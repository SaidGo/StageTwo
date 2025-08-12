-- Правки под PostgreSQL. Идемпотентно.
ALTER TABLE IF EXISTS legal_entities
    ADD COLUMN IF NOT EXISTS company_uuid UUID,
    ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;

-- Индекс для soft-delete выборок
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        WHERE c.relname = 'idx_legal_entities_deleted_at' AND n.nspname = 'public'
    ) THEN
        CREATE INDEX idx_legal_entities_deleted_at ON legal_entities (deleted_at);
    END IF;
END$$;
