CREATE TABLE bank_accounts (
    uuid UUID PRIMARY KEY,
    legal_entity_uuid UUID NOT NULL REFERENCES legal_entities(uuid) ON DELETE CASCADE,
    name TEXT NOT NULL,
    iban TEXT NOT NULL,
    bic TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
