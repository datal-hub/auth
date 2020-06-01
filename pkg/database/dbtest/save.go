package dbtest

import (
	"gopkg.in/reform.v1"

	"github.com/lib/pq"

	"github.com/datal-hub/auth/models"
)

// Save is imitation method saving reform model to database
func (db *DbTest) Save(model reform.Record) error {
	if model.(*models.Credentials).Login == db.ExistUser.Login {
		return &pq.Error{Code: "23505"}
	}
	return nil
}
