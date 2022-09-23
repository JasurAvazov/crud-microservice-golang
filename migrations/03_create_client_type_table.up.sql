CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "client_type_table" (
    "code" varchar NOT NULL PRIMARY KEY,
    "short_name" varchar NOT NULL,
    "long_name" jsonb NOT NULL,
    "status" varchar NOT NULL,
    "activated_at" timestamp with time zone
);