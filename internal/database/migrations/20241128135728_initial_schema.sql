-- +goose Up
-- +goose StatementBegin
CREATE TABLE tenant (
    id INTEGER PRIMARY KEY,
    key TEXT NOT NULL,
    name TEXT NOT NULL
);

CREATE INDEX tenant_key ON tenant (key);

CREATE TABLE prescription (
    id INTEGER PRIMARY KEY,
    tenant INTEGER NOT NULL,
    interval INTEGER NOT NULL,
    interval_unit TEXT NOT NULL,
    start_at TEXT NOT NULL,
    end_at TEXT NULL,
    medicine TEXT NOT NULL,
    amount TEXT NULL,
    FOREIGN KEY (tenant) REFERENCES tenant (id)
);

CREATE INDEX prescription_tenant ON prescription (tenant);

CREATE TABLE serving (
    id INTEGER PRIMARY KEY,
    tenant INTEGER NOT NULL,
    prescription INTEGER NOT NULL,
    occurrence INTEGER NOT NULL,
    amount TEXT NULL,
    taken BOOLEAN DEFAULT FALSE,
    taken_at TEXT NULL,
    FOREIGN KEY (tenant) REFERENCES tenant (id),
    FOREIGN KEY (prescription) REFERENCES prescription (id)
);

CREATE INDEX serving_tenant ON serving (tenant);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE serving;

DROP TABLE prescription;

DROP TABLE tenant;

-- +goose StatementEnd
