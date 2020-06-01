package logger

import (
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
