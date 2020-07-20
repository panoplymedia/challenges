CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE items(
       id uuid          PRIMARY KEY DEFAULT uuid_generate_v4(),
       description text NOT NULL UNIQUE,
       price NUMERIC
);

CREATE TABLE merchants(
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       name text NOT NULL,
       address text NOT NULL,
       UNIQUE (name, address)
);

CREATE TABLE sales (
       customer_name text,
       item_id uuid,
       quantity integer,
       merchant_id uuid,
       FOREIGN KEY (item_id) REFERENCES items (id),
       FOREIGN KEY (merchant_id) REFERENCES merchants (id)
);
