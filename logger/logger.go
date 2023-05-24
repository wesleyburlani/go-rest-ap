package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/config"
)

func NewLogger(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	if level, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logger.SetLevel(level)
	}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.AddHook(&TelemetryHook{})
	return logger
}
