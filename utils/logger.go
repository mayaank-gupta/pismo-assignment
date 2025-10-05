package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var baseLogger *logrus.Logger

func InitLogger() {
	baseLogger = logrus.New()
	logrus.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	baseLogger.SetOutput(os.Stdout)
	baseLogger.SetLevel(logrus.InfoLevel)
	baseLogger.SetFormatter(&logrus.JSONFormatter{})
}

func GetLogger(component string) *logrus.Entry {
	if baseLogger == nil {
		InitLogger()
	}
	return baseLogger.WithField("name", component)
}

func getLoggerLevel(value string) logrus.Level {
	switch value {
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
