-- +goose Up
-- +goose StatementBegin
ALTER TABLE tenant
DROP COLUMN name;

-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
ADD COLUMN name TEXT NOT NULL;

-- +goose StatementEnd
