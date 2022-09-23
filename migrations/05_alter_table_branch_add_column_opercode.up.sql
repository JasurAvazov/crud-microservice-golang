BEGIN;

ALTER TABLE "branches_table"
    ADD COLUMN "oper_code" VARCHAR;

COMMIT;