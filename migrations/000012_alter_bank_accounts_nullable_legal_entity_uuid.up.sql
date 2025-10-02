-- 000012 up
ALTER TABLE bank_accounts
  ALTER COLUMN legal_entity_uuid DROP NOT NULL;

-- снимем FK, если он есть
DO $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_name = 'bank_accounts'
      AND constraint_type = 'FOREIGN KEY'
      AND constraint_name = 'bank_accounts_legal_entity_uuid_fkey'
  ) THEN
    ALTER TABLE bank_accounts DROP CONSTRAINT bank_accounts_legal_entity_uuid_fkey;
  END IF;
END$$;

-- индекc по колонке для быстрых выборок по LE
CREATE INDEX IF NOT EXISTS idx_bank_accounts_legal_entity_uuid
  ON bank_accounts(legal_entity_uuid);
