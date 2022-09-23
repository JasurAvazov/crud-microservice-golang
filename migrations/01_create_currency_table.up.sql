CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "currency_tab" (
    "code" varchar NOT NULL PRIMARY KEY,
    "cur_code" varchar NOT NULL,
    "name" jsonb NOT NULL,
    "activated_at" timestamp with time zone
);