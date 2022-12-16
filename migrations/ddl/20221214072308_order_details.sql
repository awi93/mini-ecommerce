-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_details (
    id bigserial primary key,
    order_id bigint not null,
    product_id bigint not null,
    quantity int not null,
    unit_price numeric(18,4) not null,
    discount numeric(18,4) not null,
    foreign key (order_id) references orders (id) on update cascade on delete cascade,
    foreign key (product_id) references products (id) on update cascade on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_details;
-- +goose StatementEnd
