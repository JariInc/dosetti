-- +goose Up
ALTER TABLE tenant
DROP COLUMN name;

-- +goose Down
ALTER TABLE
ADD COLUMN name TEXT NOT NULL;