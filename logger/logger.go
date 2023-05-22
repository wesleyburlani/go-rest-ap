package logger

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.AddHook(&TelemetryHook{})
	return logger
}
