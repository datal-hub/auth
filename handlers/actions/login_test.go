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
)

func TestLoginOK(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Login OK",
		ExpectedCode: http.StatusOK,
	}
	existUserJson, _ := json.Marshal(testData.ExistUser)
	body := strings.NewReader(string(existUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/login", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Login(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusOK, resp.Code, test.Description)
}

func TestLoginNotExist(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Login not exist",
		ExpectedCode: http.StatusNotFound,
	}
	newUserJson, _ := json.Marshal(testData.NewUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/login", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Login(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusNotFound, resp.Code, test.Description)
}

func TestLoginEmpty(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Login empty",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyLoginUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/login", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Login(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}

func TestLoginEmptyPassword(t *testing.T) {
	test := middleware.TestDescription{
		Description:  "Login empty password",
		ExpectedCode: http.StatusBadRequest,
	}
	newUserJson, _ := json.Marshal(testData.EmptyPasswordUser)
	body := strings.NewReader(string(newUserJson))
	var req = httptest.NewRequest("POST", "http://auth.ru/login", body)
	resp := httptest.NewRecorder()
	tCtx := ctx()
	req = tCtx.InitContext(req)
	Login(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code, test.Description)
}
