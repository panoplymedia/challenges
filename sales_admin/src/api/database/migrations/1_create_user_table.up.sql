CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS "user"
(
    id            int unsigned auto_increment primary key,
    username      varchar(255) not null,
    password_hash varchar(255) not null,
    email         varchar(255) not null,
    blocked       bool         not null default false,
    date_blocked  timestamp             default null,
    deleted_at    timestamp             default null,
    created_at    timestamp    not null default current_date,
    updated_at    timestamp    not null default current_date
);