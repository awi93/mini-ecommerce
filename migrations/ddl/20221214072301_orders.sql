-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    id bigserial primary key,
    customer_id bigint not null,
    employee_id bigint not null,
    order_date date,
    purchase_order_number varchar(30) not null,
    ship_date date,
    shipping_method_id bigint not null,
    freight_charge numeric(18,4) not null,
    taxes numeric(18,4) not null,
    payment_received bool not null default false,
    comment varchar(150) null,
    foreign key (customer_id) references customers(id) on update cascade on delete cascade,
    foreign key (employee_id) references employees(id) on update cascade on delete cascade,
    foreign key (shipping_method_id) references shipping_methods(id) on update cascade on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
