-- +migrate Up
CREATE TABLE IF NOT EXISTS "user" (
    id bigserial NOT NULL,
    username varchar NOT NULL,
    "password" varchar NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (id),
    CONSTRAINT user_un UNIQUE (username)
);

-- +migrate Down