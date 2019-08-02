CREATE TABLE IF NOT EXISTS sale
(
    id          serial primary key,
    deleted_at  timestamp          default null,
    created_at  timestamp not null default current_date,
    updated_at  timestamp not null default current_date,
    customer_id integer   not null references customer (id),
    product_id  integer   not null references product (id),
    merchant_id integer   not null references merchant (id),
    quantity    integer   not null,
    total_price float     not null
);