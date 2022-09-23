BEGIN;

ALTER TABLE "branches_table"
    ADD COLUMN "coords" jsonb not null default '{}'::jsonb;
COMMIT;