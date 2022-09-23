CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "maritials_status" (
    "code" varchar NOT NULL PRIMARY KEY,
    "gender" jsonb NOT NULL,
    "status" jsonb NOT NULL,
    "activated_at" timestamp with time zone
);