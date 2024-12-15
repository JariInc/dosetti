-- +goose Up
-- +goose StatementBegin
CREATE TABLE tenant (id INTEGER PRIMARY KEY, key TEXT NOT NULL) RANDOM ROWID;

CREATE TABLE prescription (
    id INTEGER PRIMARY KEY,
    tenant INTEGER NOT NULL,
    interval INTEGER NOT NULL,
    interval_unit TEXT NOT NULL,
    start_date TEXT NOT NULL,
    "offset" INTEGER DEFAULT 0,
    medicine TEXT NOT NULL,
    amount TEXT NULL,
    FOREIGN KEY (tenant) REFERENCES tenant (id)
) RANDOM ROWID;

CREATE INDEX prescription_tenant ON prescription (tenant);

CREATE TABLE serving (
    id INTEGER PRIMARY KEY,
    tenant INTEGER NOT NULL,
    prescription INTEGER NOT NULL,
    medicine TEXT NOT NULL,
    amount TEXT NULL,
    taken BOOLEAN DEFAULT FALSE,
    scheduled_at TEXT NOT NULL,
    taken_at TEXT NOT NULL,
    FOREIGN KEY (tenant) REFERENCES tenant (id),
    FOREIGN KEY (prescription) REFERENCES prescription (id)
) RANDOM ROWID;

CREATE INDEX serving_tenant ON serving (tenant);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE serving;

DROP TABLE prescription;

DROP TABLE tenant;

-- +goose StatementEnd
