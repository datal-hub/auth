package models_test

import (
	"errors"
	"testing"

	. "github.com/datal-hub/auth/models"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		Cred Credentials
		Err  error
	}{
		{
			Cred: Credentials{Login: "test", Email: "test@test.com", Phone: "+79998887766", Password: "test12345"},
			Err:  nil},
		{
			Cred: Credentials{Login: "", Email: "test@test.com", Phone: "+79998887766", Password: "test12345"},
			Err:  errors.New("empty field"),
		},
		{
			Cred: Credentials{Login: "test", Email: "", Phone: "+79998887766", Password: "test12345"},
			Err:  errors.New("empty field"),
		},
		{
			Cred: Credentials{Login: "test", Email: "test@test.com", Phone: "", Password: "test12345"},
			Err:  errors.New("empty field"),
		},
		{
			Cred: Credentials{Login: "test", Email: "test@test.com", Phone: "+79998887766", Password: ""},
			Err:  errors.New("empty field"),
		},
		{
			Cred: Credentials{Login: "test", Email: "invalid", Phone: "+79998887766", Password: "test12345"},
			Err:  errors.New("invalid email format"),
		},
		{
			Cred: Credentials{Login: "test", Email: "test@test.com", Phone: "12121212", Password: "test12345"},
			Err:  errors.New("invalid phone format"),
		},
	}
	for i, tc := range testCases {
		if err := tc.Cred.IsValid(); (err == nil && err != tc.Err) || (err != nil && err.Error() != tc.Err.Error()) {
			t.Fatalf("Failed test number: %d", i)
		}
	}
}
