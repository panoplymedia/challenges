CREATE TABLE IF NOT EXISTS customer
(
    id         serial       not null primary key,
    deleted_at timestamp             default null,
    created_at timestamp    not null default current_date,
    updated_at timestamp    not null default current_date,
    full_name  varchar(255) not null,
    email      varchar(255) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_customer_email on customer(email);