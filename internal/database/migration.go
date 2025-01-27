package database

import (
	"context"
	"database/sql"

	goose "github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/go-libsql"
)

func Migrate(ctx context.Context, connection *sql.DB) error {
	if err := goose.SetDialect("turso"); err != nil {
		return err
	}

	return goose.UpContext(ctx, connection, "./migrations")
}
