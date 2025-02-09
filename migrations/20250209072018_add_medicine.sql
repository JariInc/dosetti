-- +goose Up
CREATE TABLE medicine (
    id INTEGER PRIMARY KEY,
    tenant INTEGER NOT NULL,
    name TEXT NULL,
    doses_left REAL NULL,
    FOREIGN KEY (tenant) REFERENCES tenant (id)
);

CREATE INDEX medicine_tenant ON medicine (tenant);

ALTER TABLE prescription DROP COLUMN medicine;
ALTER TABLE prescription ADD COLUMN medicine INTEGER NULL;

ALTER TABLE prescription DROP COLUMN amount;
ALTER TABLE prescription ADD COLUMN amount REAL NULL;

ALTER TABLE serving DROP COLUMN amount;
ALTER TABLE serving ADD COLUMN amount REAL NULL;

-- +goose Down
DROP TABLE medicine;
ALTER TABLE prescription DROP COLUMN medicine;
ALTER TABLE prescription ADD COLUMN medicine TEXT DEFAULT "medicine" NOT NULL;

ALTER TABLE prescription DROP COLUMN amount;
ALTER TABLE prescription ADD COLUMN amount TEXT NULL;

ALTER TABLE serving DROP COLUMN amount;
ALTER TABLE serving ADD COLUMN amount TEXT NULL;
