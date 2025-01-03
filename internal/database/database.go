package database

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

const (
	DATE_TIME_FORMAT = "2006-01-02T15:04:05.999999999Z07:00"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Close() {
	db.Conn.Close()
}

func NewDatabase(connectionURL string) (*Database, error) {
	conn, err := sql.Open("libsql", connectionURL)

	if err != nil {
		return nil, err
	}

	db := &Database{
		Conn: conn,
	}

	return db, nil
}
