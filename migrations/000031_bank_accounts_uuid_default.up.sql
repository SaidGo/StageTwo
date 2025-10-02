CREATE EXTENSION IF NOT EXISTS pgcrypto;
ALTER TABLE bank_accounts
    ALTER COLUMN uuid SET DEFAULT gen_random_uuid();
