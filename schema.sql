CREATE DATABASE achs;
USE achs;
CREATE TABLE salesdata (
  id SERIAL PRIMARY KEY,
  customer_name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  price decimal(20) NOT NULL,
  quantity int NOT NULL,
  merchant_name varchar(255) NOT NULL,
  merchant_address varchar(255) NOT NULL
);