-- +goose Up
-- +goose StatementBegin
ALTER TABLE tenant
RENAME COLUMN key TO uuid;

ALTER TABLE tenant
DROP COLUMN name;

-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
ALTER TABLE tenant
RENAME COLUMN uuid TO key;

ALTER TABLE
ADD COLUMN name TEXT NOT NULL;

-- +goose StatementEnd
