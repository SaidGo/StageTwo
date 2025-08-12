-- Аккуратный откат (если столбцы существуют)
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.columns
               WHERE table_name='legal_entities' AND column_name='deleted_at') THEN
        ALTER TABLE legal_entities DROP COLUMN deleted_at;
    END IF;
    IF EXISTS (SELECT 1 FROM information_schema.columns
               WHERE table_name='legal_entities' AND column_name='company_uuid') THEN
        ALTER TABLE legal_entities DROP COLUMN company_uuid;
    END IF;
END$$;

DROP INDEX IF EXISTS idx_legal_entities_deleted_at;
