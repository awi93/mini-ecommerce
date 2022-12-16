-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id bigserial primary key,
    product_name varchar(255) not null,
    unit_price numeric(19,4) not null,
    in_stock bool not null default true
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
