CREATE TABLE users (
    uuid TEXT PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    lname TEXT NOT NULL DEFAULT '',
    pname TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    phone INTEGER NOT NULL DEFAULT 0,
    is_valid BOOLEAN NOT NULL DEFAULT 0,
    password TEXT NOT NULL DEFAULT '',
    provider INTEGER NOT NULL DEFAULT 0,
    color TEXT NOT NULL DEFAULT '#000000',
    has_photo BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    valid_at DATETIME,
    validation_send_at DATETIME,
    reset_send_at DATETIME,
    meta TEXT NOT NULL DEFAULT '{}',
    position TEXT NOT NULL DEFAULT '{}'
);

CREATE UNIQUE INDEX users_email_unique ON users (email);
CREATE INDEX users_created_at ON users (created_at DESC);
CREATE INDEX users_updated_at ON users (updated_at DESC);
