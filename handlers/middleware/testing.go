package middleware

import (
	"net/http"

	"github.com/datal-hub/auth/pkg/database"
)

type TestDescription struct {
	Description  string
	Url          string
	QueryValues  map[string]string
	ExpectedBody string
	ExpectedCode int
}

type TestContext struct {
	DB database.DB
}

func (tCtx TestContext) InitContext(r *http.Request) *http.Request {
	r = SetContext(r, tCtx.DB)
	return r
}

func (tCtx TestContext) InitContextHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = tCtx.InitContext(r)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
