package http_middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BeforeRequestLoggerMiddleware struct {
	logger *logrus.Logger
}

func NewBeforeRequestLoggerMiddleware(logger *logrus.Logger) *BeforeRequestLoggerMiddleware {
	return &BeforeRequestLoggerMiddleware{
		logger,
	}
}

func (instance *BeforeRequestLoggerMiddleware) Handle(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	path := c.Request.URL.Path
	clientIP := c.ClientIP()
	clientUserAgent := c.Request.UserAgent()
	dataLength := c.Writer.Size()
	if dataLength < 0 {
		dataLength = 0
	}

	entry := instance.logger.WithFields(logrus.Fields{
		"hostname":  hostname,
		"clientIP":  clientIP,
		"method":    c.Request.Method,
		"path":      path,
		"userAgent": clientUserAgent,
	})
	entry.Debug("request received")
	c.Next()
}
