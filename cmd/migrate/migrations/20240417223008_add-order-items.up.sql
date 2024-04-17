create table if not exists order_items (
    id integer not null primary key,
    order_id integer not null,
    product_id integer not null,
    quantity integer not null,
    created_at timestamp default current_timestamp not null,
    foreign key (order_id) references orders (id),
    foreign key (product_id) references products (id)
);
