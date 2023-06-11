package logger

import (
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	LogLevel string
}

func NewLogger(config *LoggerConfig) *logrus.Logger {
	logger := logrus.New()

	if level, err := logrus.ParseLevel(config.LogLevel); err == nil {
		logger.SetLevel(level)
	}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.AddHook(&TelemetryHook{})
	return logger
}
