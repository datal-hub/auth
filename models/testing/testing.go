package testing

import . "github.com/datal-hub/auth/models"

var NewUser = Credentials{
	Login:    "user",
	Phone:    "+79998887766",
	Email:    "test@test.ru",
	Password: "testpassword",
}

var ExistUser = Credentials{
	Login:    "exist",
	Phone:    "+79998887766",
	Email:    "test@test.ru",
	Password: "testpassword",
}

var EmptyLoginUser = Credentials{
	Login:    "",
	Phone:    "+79998887766",
	Email:    "test@test.ru",
	Password: "testpassword",
}

var EmptyPhoneUser = Credentials{
	Login:    "user",
	Phone:    "",
	Email:    "test@test.ru",
	Password: "testpassword",
}

var EmptyEmailUser = Credentials{
	Login:    "user",
	Phone:    "+79998887766",
	Email:    "",
	Password: "testpassword",
}

var EmptyPasswordUser = Credentials{
	Login:    "user",
	Phone:    "+79998887766",
	Email:    "test@test.ru",
	Password: "",
}
