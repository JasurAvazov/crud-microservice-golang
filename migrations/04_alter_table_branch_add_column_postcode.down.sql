BEGIN;

ALTER TABLE "branches_table"
    DROP COLUMN "postal_code";

COMMIT;