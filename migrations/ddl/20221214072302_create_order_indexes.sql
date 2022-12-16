-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX order_purchase_order_number_idx ON orders (purchase_order_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX order_purchase_order_number_idx;
-- +goose StatementEnd
