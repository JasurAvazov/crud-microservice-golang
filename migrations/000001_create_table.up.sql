CREATE TABLE IF NOT EXISTS countries (
        "code" varchar NOT NULL UNIQUE ,
        "title" jsonb NOT NULL,
        "alpha_duo" varchar,
        "alpha_triple" varchar,
        "currency_code" varchar,
        "location_sign" varchar,
        "activated_at" timestamp with time zone
);


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

CREATE TABLE IF NOT EXISTS regions (
        "code" varchar NOT NULL PRIMARY KEY,
        "title" jsonb NOT NULL,
        "rank" varchar,
        "activated_at" timestamp with time zone,
        "deactivated_at" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS nationalities (
        "code" varchar NOT NULL PRIMARY KEY,
        "title" jsonb NOT NULL,
        "activated_at" timestamp with time zone,
        "deactivated_at" timestamp with time zone
);