package pgsql

import (
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/datal-hub/auth/models"
)

// GetCredentials is method finding credentials in database
func (db *PgSQL) GetCredentials(login string) (*models.Credentials, error) {
	rdb := reform.NewDB(db.SqlDB(), postgresql.Dialect, nil)
	var credentials models.Credentials
	err := rdb.FindOneTo(&credentials, "login", login)
	if err != nil && err == reform.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &credentials, nil
}
