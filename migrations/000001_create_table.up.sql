CREATE TABLE IF NOT EXISTS customer (
        "id" int NOT NULL,
        "name" varchar(14) NOT NULL,
        "country" char(3) NOT NULL,
        "address" text,
        "phone" varchar(50),
        PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS "order" (
        "id" int NOT NULL,
        "date" date NOT NULL,
        "cust_id" int NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (cust_id) REFERENCES customer(id)
);

CREATE TABLE IF NOT EXISTS category (
        "id" int NOT NULL,
        "name" varchar(250),
        PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product (
        "id" int NOT NULL,
        "name" varchar(10),
        "category_id" int NOT NULL,
        "description" varchar(20),
        "price" numeric(6,2),
        "photo" varchar(1024),
        PRIMARY KEY (id),
        FOREIGN KEY (category_id) REFERENCES category(id)
);