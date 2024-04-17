create table if not exists products (
    id integer not null primary key,
    name text not null,
    description text not null,
    image text not null,
    quantity integer not null,
    created_at timestamp default current_timestamp not null
)
