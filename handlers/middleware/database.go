package middleware

import (
	"context"
	"net/http"

	. "github.com/datal-hub/auth/pkg/database"
	log "github.com/datal-hub/auth/pkg/logger"
)

type contextKey int

const dbContextKey contextKey = 0

// Retrieves database connection from the context.
func DBFromContext(ctx context.Context) DB {
	db, ok := ctx.Value(dbContextKey).(DB)
	if !ok {
		panic("database.FromContext: no database in the context. WithDatabaseHandler must be in the handlers chain.")
	}
	return db
}

func SetContext(r *http.Request, db DB) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), dbContextKey, db))
}

// Creates database connection and stores it in context, so every
// handler in the chain could use it.
func DatabaseHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Middleware: WithDatabase")
		db, err := NewDB()
		if err != nil {
			log.ErrorF("WithDatabaseHandler: error creating database connection",
				log.Fields{"message": err})
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"message":"internal error"}`))
			return
		}
		defer func() {
			log.Debug("Closing database connection.")
			if err := db.Close(); err != nil {
				log.Error(err.Error())
			}
		}()
		r = SetContext(r, db)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
