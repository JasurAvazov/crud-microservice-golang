CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "branches_table" (
    "code" varchar NOT NULL PRIMARY KEY,
    "title" jsonb NOT NULL,
    "mfo" varchar NOT NULL,
    "address" jsonb,
    "activated_at" timestamp with time zone
);