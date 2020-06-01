package actions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datal-hub/auth/handlers/middleware"
	testData "github.com/datal-hub/auth/models/testing"
	"github.com/datal-hub/auth/pkg/database"
	"github.com/datal-hub/auth/pkg/settings"
)

func ctx() middleware.TestContext {
	var tCtx middleware.TestContext
	settings.VerboseMode = true
	database.Testing = true
	tCtx.DB, _ = database.NewDB()
	return tCtx
}

func TestRegisterOK(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register OK",
		ExpectedCode: http.StatusOK,
	}
	newUserJson, _ := json.Marshal(testData.NewUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusOK, resp.Code, test.Description)
}

func TestRegisterExistUser(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register exist user",
		ExpectedCode: http.StatusConflict,
	}
	newUserJson, _ := json.Marshal(testData.ExistUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusConflict, resp.Code, test.Description)
}

func TestRegisterEmptyLogin(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register empty login",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyLoginUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}

func TestRegisterEmptyEmail(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register empty email",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyEmailUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}

func TestRegisterEmptyPassword(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register empty password",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyPasswordUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}

func TestRegisterEmptyPhone(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Register empty phone",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyPhoneUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/register", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Register(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}
