package models

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//go:generate reform

// Credentials is the model desired registration and authentication data
//reform:credentials
type Credentials struct {
	ID         int64     `reform:"id,pk" json:"-"`
	Login      string    `reform:"login" json:"login"`
	Email      string    `reform:"email" json:"email"`
	Phone      string    `reform:"phone" json:"phone"`
	Password   string    `json:"password"`
	Hash       []byte    `reform:"password" json:"-"`
	CreateDttm time.Time `reform:"create_dttm" json:"-"`
}

var (
	phoneRegexp = regexp.MustCompile(`((\+7)+([0-9]){10})`)
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (c *Credentials) IsValid() error {
	if c.Login == "" || c.Phone == "" || c.Email == "" || c.Password == "" {
		return errors.New("empty field")
	}
	if c.Phone != "" && !phoneRegexp.MatchString(c.Phone) {
		return errors.New("invalid phone format")
	}
	if c.Email != "" && !emailRegexp.MatchString(c.Email) {
		return errors.New("invalid email format")
	}
	if len(c.Password) < 8 {
		return errors.New("password is too short")
	}
	return nil
}

func (c *Credentials) SetHash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Hash = hashedPassword
	return nil
}

func (c Credentials) CheckPassword() error {
	return bcrypt.CompareHashAndPassword(c.Hash, []byte(c.Password))
}
