\echo '== Enable pgcrypto, drop FK on legal_entity_uuid, make it NULLABLE, set uuid default'
CREATE EXTENSION IF NOT EXISTS pgcrypto;

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

ALTER TABLE bank_accounts
    ALTER COLUMN legal_entity_uuid DROP NOT NULL;

ALTER TABLE bank_accounts
    ALTER COLUMN uuid SET DEFAULT gen_random_uuid();
\echo '== Done'
