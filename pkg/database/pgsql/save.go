package pgsql

import (
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

// Save is method saving reform model to database
func (db *PgSQL) Save(model reform.Record) error {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	if err := rdb.Save(model); err != nil {
		return err
	}
	return nil
}
