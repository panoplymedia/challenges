CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "user"
(
    id         serial       not null primary key,
    deleted_at timestamp             default null,
    created_at timestamp    not null default current_date,
    updated_at timestamp    not null default current_date,
    password   varchar(255) not null,
    email      varchar(255) not null,
    role       varchar(255) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email on "user" (email);