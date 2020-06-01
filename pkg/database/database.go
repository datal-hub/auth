package database

import (
	"database/sql"

	"gopkg.in/reform.v1"

	_ "github.com/lib/pq"

	"github.com/datal-hub/auth/models"
	"github.com/datal-hub/auth/pkg/database/pgsql"
	"github.com/datal-hub/auth/pkg/settings"
)

// DB define interface for work with database
type DB interface {
	IsEmpty() bool
	Clear()
	Init(force bool) error
	Close() error
	SqlDB() *sql.DB

	Save(model reform.Record) error
	GetCredentials(login string) (*models.Credentials, error)
}

// NewDB create new database connection accordance with mode
func NewDB() (DB, error) {

	if Testing {
		return testDb()
	}

	// postgresql: //[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	db, err := sql.Open("postgres", settings.DB.Url)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &pgsql.PgSQL{DB: db}, nil
}
