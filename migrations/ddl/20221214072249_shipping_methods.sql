-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shipping_methods (
    id bigserial primary key,
    shipping_method varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shipping_methods;
-- +goose StatementEnd
