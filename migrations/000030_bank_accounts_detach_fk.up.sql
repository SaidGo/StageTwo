CREATE TABLE IF NOT EXISTS bank_accounts (
    uuid               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    legal_entity_uuid  UUID NULL,
    bik                TEXT,
    bank               TEXT,
    address            TEXT,
    corr_account       TEXT,
    account            TEXT NOT NULL,
    currency           TEXT,
    comment            TEXT,
    is_primary         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at         TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at         TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_bank_accounts_legal_entity_uuid ON bank_accounts(legal_entity_uuid);
CREATE INDEX IF NOT EXISTS idx_bank_accounts_deleted_at ON bank_accounts(deleted_at);

ALTER TABLE bank_accounts ALTER COLUMN legal_entity_uuid DROP NOT NULL;

DO $$
BEGIN
  ALTER TABLE bank_accounts DROP CONSTRAINT bank_accounts_legal_entity_uuid_fkey;
EXCEPTION WHEN undefined_object THEN
  NULL;
END $$;
