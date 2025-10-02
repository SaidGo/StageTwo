-- 000012 down (возврат обратно: колонка NOT NULL + FK)
ALTER TABLE bank_accounts
  ALTER COLUMN legal_entity_uuid SET NOT NULL;

ALTER TABLE bank_accounts
  ADD CONSTRAINT bank_accounts_legal_entity_uuid_fkey
  FOREIGN KEY (legal_entity_uuid)
  REFERENCES legal_entities (uuid);
