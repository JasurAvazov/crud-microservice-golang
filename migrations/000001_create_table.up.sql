CREATE TABLE IF NOT EXISTS customer (
        "id" int NOT NULL,
        "name" varchar(14) NOT NULL,
        "country" char(3) NOT NULL,
        "address" text,
        "phone" varchar(50),
        PRIMARY KEY (id)
);