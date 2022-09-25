
CREATE TABLE IF NOT EXISTS districts (
        "code" varchar NOT NULL PRIMARY KEY,
        "soato_code" varchar,
        "gni_code" varchar,
        "province_code" varchar,
        "title" jsonb NOT NULL,
        "activated_at" timestamp with time zone,
        "deactivated_at" timestamp with time zone,
        "state" bool default true
);

CREATE TABLE IF NOT EXISTS customer (
        "id" integer NOT NULL,
        "name" varchar(14) NOT NULL,
        "country" char(3) NOT NULL,
        "address" text,
        "phone" varchar(50),
        PRIMARY KEY (id)
);