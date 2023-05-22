package http_middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type LoggerMiddleware struct {
	logger *logrus.Logger
}

func NewLoggerMiddleware(logger *logrus.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger,
	}
}

func (instance *LoggerMiddleware) Handle(c *gin.Context) {
	ginlogrus.Logger(instance.logger.WithContext(c.Request.Context()))(c)
}
