BEGIN;

ALTER TABLE "branches_table"
    DROP COLUMN "oper_code";

COMMIT;