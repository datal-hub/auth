package middleware

import (
	"net/http"
	"time"

	. "github.com/datal-hub/auth/pkg/logger"
)

func HttpLogDetails(r *http.Request) Fields {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return Fields{
		"IP":     ip,
		"URL":    r.URL.String(),
		"Method": r.Method,
	}
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fields := HttpLogDetails(r)
		DebugF("Route handler start.", fields)
		dttmStart := time.Now()
		next.ServeHTTP(w, r)
		dttmStop := time.Now()
		fields["elapsed"] = dttmStop.Sub(dttmStart)
		DebugF("Route handler stop.", fields)
	}
	return http.HandlerFunc(fn)
}
