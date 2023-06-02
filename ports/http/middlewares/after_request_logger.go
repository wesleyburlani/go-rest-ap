package middlewares

import (
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type AfterRequestLoggerMiddleware struct {
	logger *logrus.Logger
}

func NewAfterRquestLoggerMiddleware(logger *logrus.Logger) *AfterRequestLoggerMiddleware {
	return &AfterRequestLoggerMiddleware{
		logger,
	}
}

func (instance *AfterRequestLoggerMiddleware) Handle(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	path := c.Request.URL.Path
	start := time.Now()
	c.Next()
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1_000_000.0))
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
		"hostname":   hostname,
		"statusCode": statusCode,
		"latency":    latency, // time to process
		"clientIP":   clientIP,
		"method":     c.Request.Method,
		"path":       path,
		"referer":    referer,
		"dataLength": dataLength,
		"userAgent":  clientUserAgent,
		"traceId":    traceId,
		"spanId":     spanId,
	})

	if len(c.Errors) > 0 {
		entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		msg := "response sent"
		if statusCode >= http.StatusInternalServerError {
			entry.Error(msg)
		} else if statusCode >= http.StatusBadRequest {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}
	}
}
