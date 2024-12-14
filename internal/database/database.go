package database

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
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
