package libsql

import (
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

const (
	DATE_TIME_FORMAT = "2006-01-02T15:04:05.999999999Z07:00"
)

type Connection struct {
	Conn *sql.DB
}

func (db *Connection) Close() {
	db.Conn.Close()
}

func NewConnection(connectionURL string) (*Connection, error) {
	conn, err := sql.Open("libsql", connectionURL)

	if err != nil {
		return nil, err
	}

	db := &Connection{
		Conn: conn,
	}

	return db, nil
}
