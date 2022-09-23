CREATE TABLE IF NOT EXISTS "offers" (
    "code" char(5) NOT NULL PRIMARY KEY,
    "name" jsonb NOT NULL,
    "should_be_signed" boolean NOT NULL DEFAULT FALSE,
    "signature_type" varchar,
    "created_at" timestamptz default current_timestamp
);