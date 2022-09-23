BEGIN;

ALTER TABLE "branches_table"
    DROP COLUMN "coords";

COMMIT;