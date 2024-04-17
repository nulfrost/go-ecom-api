create table if not exists orders (
    id integer not null primary key,
    user_id integer not null,
    total integer not null,
    status text not null,
    address text not null,
    created_at timestamp default current_timestamp not null,
    foreign key (user_id) references users (id)
)
