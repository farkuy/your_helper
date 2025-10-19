-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN some_table TEXT DEFAULT 'some text';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN some_table;
-- +goose StatementEnd