package dbtest

import (
	"github.com/datal-hub/auth/models"
)

func (db *DbTest) GetCredentials(login string) (*models.Credentials, error) {
	if db.ExistUser.Login == login {
		res := &db.ExistUser
		res.Hash = []byte("$2a$10$x.B6UdBeQ2lF2pwXQjpeyOvk3h6j/iFnWrU.4uuhSsbslR0Xrw71S")
		return res, nil
	}
	return nil, nil
}
