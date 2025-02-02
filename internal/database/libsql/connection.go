package libsql

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

const (
	DATE_TIME_FORMAT = "2006-01-02T15:04:05.999999999Z07:00"
)

func NewConnection(connectionURL string) (*sql.DB, error) {
	conn, err := sql.Open("libsql", connectionURL)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
