CREATE TABLE IF NOT EXISTS product
(
    id          serial primary key,
    deleted_at  timestamp             default null,
    created_at  timestamp    not null default current_date,
    updated_at  timestamp    not null default current_date,
    merchant_id bigint       not null references merchant (id),
    name        varchar(255) not null,
    description varchar(255)          default null,
    item_price  float        not null
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_product_name_merchant ON product(name, merchant_id);