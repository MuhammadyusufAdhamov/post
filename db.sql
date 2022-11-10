create table posts(
    id serial primary key,
    post varchar,
    user_id numeric,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);