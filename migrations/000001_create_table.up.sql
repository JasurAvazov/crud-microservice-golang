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

CREATE TABLE IF NOT EXISTS detail (
        "id" int NOT NULL,
        "ord_id" int NOT NULL,
        "pr_id" int NOT NULL,
        "quantity" int NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (ord_id) REFERENCES "order"(id),
        FOREIGN KEY (pr_id) REFERENCES  product(id)
);

CREATE TABLE IF NOT EXISTS invoice (
        "id" int NOT NULL,
        "ord_id" int NOT NULL,
        "amount" numeric(8,2) NOT NULL,
        "issued" date NOT NULL,
        "due" date NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (ord_id) REFERENCES "order"(id)
);

CREATE TABLE IF NOT EXISTS payment (
        "id" int NOT NULL,
        "time" timestamp with time zone NOT NULL,
        "amount" numeric(8,2) NOT NULL,
        "inv_id" int NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (inv_id) REFERENCES invoice(id)
);