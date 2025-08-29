ALTER TABLE bank_accounts
    ALTER COLUMN uuid DROP DEFAULT;

ALTER TABLE bank_accounts
    ALTER COLUMN legal_entity_uuid SET NOT NULL;
