package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
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

	span := trace.SpanFromContext(c.Request.Context())
	spanContext := span.SpanContext()
	traceId := spanContext.TraceID()
	spanId := spanContext.SpanID()

	entry := instance.logger.WithFields(logrus.Fields{
		"hostname":  hostname,
		"clientIP":  clientIP,
		"method":    c.Request.Method,
		"path":      path,
		"userAgent": clientUserAgent,
		"traceId":   traceId,
		"spanId":    spanId,
	})
	entry.Debug("request received")
	c.Next()
}
