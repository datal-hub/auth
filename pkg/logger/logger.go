package logger

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/datal-hub/auth/pkg/settings"
)

type Fields map[string]interface{}

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if settings.VerboseMode {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func Debug(msg string) {
	logrus.Debug(msg)
}

func Info(msg string) {
	logrus.Info(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}

func DebugF(msg string, fields Fields) {
	logrusFields := logrus.Fields(fields)
	logrus.WithFields(logrusFields).Debug(msg)
}

func InfoF(msg string, fields Fields) {
	logrusFields := logrus.Fields(fields)
	logrus.WithFields(logrusFields).Info(msg)
}

func ErrorF(msg string, fields Fields) {
	logrusFields := logrus.Fields(fields)
	logrus.WithFields(logrusFields).Error(msg)
}

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
