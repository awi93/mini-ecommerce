-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees (
    id bigserial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) null,
    title varchar(255) not null,
    work_phone varchar(30) null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
