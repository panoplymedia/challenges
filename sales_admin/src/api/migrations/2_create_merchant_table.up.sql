CREATE TABLE IF NOT EXISTS merchant
(
    id         serial       not null primary key,
    deleted_at timestamp             default null,
    created_at timestamp    not null default current_date,
    updated_at timestamp    not null default current_date,
    name       varchar(255) not null,
    address    varchar(255) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_merchant_name on merchant (name);