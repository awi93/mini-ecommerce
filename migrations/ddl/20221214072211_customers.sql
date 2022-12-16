-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id bigserial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) null,
    email varchar(255) null,
    phone_number varchar(30) null,
    fax_number varchar(30) null,
    billing_address varchar(500) null,
    billing_city varchar(255) null,
    billing_state varchar(255) null,
    billing_zip_code varchar(20) null,
    company_name varchar(50) null,
    company_website varchar(255) null,
    ship_address varchar(500) null,
    ship_city varchar(255) null,
    ship_state varchar(255) null,
    ship_zip_code varchar(20) null,
    ship_phone_number varchar(50) null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
