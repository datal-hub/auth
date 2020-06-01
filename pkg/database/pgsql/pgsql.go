package pgsql

import (
	"database/sql"
)

// PgSQL define environment for work with PostgreSQL database
type PgSQL struct {
	DB *sql.DB
}

// SqlDB return real PostgreSQL sql.DB
func (db *PgSQL) SqlDB() *sql.DB {
	return db.DB
}

// Close is method for close connection
func (db *PgSQL) Close() error {
	return db.DB.Close()
}
