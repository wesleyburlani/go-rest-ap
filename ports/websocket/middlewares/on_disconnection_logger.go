package middlewares

import (
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type OnDisconnectionLoggerMiddleware struct {
	logger *logrus.Logger
}

func NewOnDisconnectionLoggerMiddleware(logger *logrus.Logger) *OnDisconnectionLoggerMiddleware {
	return &OnDisconnectionLoggerMiddleware{
		logger,
	}
}

func (instance *OnDisconnectionLoggerMiddleware) Handle(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	path := c.Request.URL.Path
	start := time.Now()
	c.Next()
	stop := time.Since(start)
	timeConnected := int(math.Ceil(float64(stop.Nanoseconds()) / 1_000_000.0))
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()
	clientUserAgent := c.Request.UserAgent()
	referer := c.Request.Referer()
	dataLength := math.Max(0, float64(c.Writer.Size()))

	span := trace.SpanFromContext(c.Request.Context())
	spanContext := span.SpanContext()
	traceId := spanContext.TraceID()
	spanId := spanContext.SpanID()

	entry := instance.logger.WithFields(logrus.Fields{
		"hostname":      hostname,
		"statusCode":    statusCode,
		"timeConnected": timeConnected,
		"clientIP":      clientIP,
		"path":          path,
		"referer":       referer,
		"dataLength":    dataLength,
		"userAgent":     clientUserAgent,
		"traceId":       traceId,
		"spanId":        spanId,
	})
	entry.Debug("websocket connection closed")
}
