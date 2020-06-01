package dbtest

import (
	"database/sql"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/datal-hub/auth/models"
)

// DbTest define environment for work with mock database
type DbTest struct {
	DB        *sql.DB
	Mock      sqlmock.Sqlmock
	ExistUser models.Credentials
}

// SqlDB return imitation sql.DB
func (db *DbTest) SqlDB() *sql.DB {
	return db.DB
}

// Close simulates close connection
func (db *DbTest) Close() error {
	return db.DB.Close()
}
