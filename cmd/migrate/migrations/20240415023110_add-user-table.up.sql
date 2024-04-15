create table if not exists users (
    id integer not null primary key,
    first_name text not null,
    last_name text not null,
    email text not null unique,
    password text not null,
    created_at timestamp default current_timestamp not null
);
