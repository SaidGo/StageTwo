-- 1) Расширение для gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 2) Снять возможный FK (имя может отличаться — снимем все, что ссылается на legal_entity_uuid)
DO $$
DECLARE
    cons RECORD;
BEGIN
    FOR cons IN
        SELECT tc.constraint_name
        FROM information_schema.table_constraints tc
        JOIN information_schema.key_column_usage kcu
          ON tc.constraint_name = kcu.constraint_name
        WHERE tc.table_name = 'bank_accounts'
          AND tc.constraint_type = 'FOREIGN KEY'
          AND kcu.column_name = 'legal_entity_uuid'
    LOOP
        EXECUTE format('ALTER TABLE bank_accounts DROP CONSTRAINT %I', cons.constraint_name);
    END LOOP;
END $$;

-- 3) Сделать колонку nullable (если сейчас NOT NULL)
ALTER TABLE bank_accounts
    ALTER COLUMN legal_entity_uuid DROP NOT NULL;

-- 4) Дефолт для первичного ключа uuid (если его нет)
ALTER TABLE bank_accounts
    ALTER COLUMN uuid SET DEFAULT gen_random_uuid();
