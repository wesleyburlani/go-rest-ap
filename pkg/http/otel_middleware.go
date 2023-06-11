package http

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type OtelMiddlewareConfig struct {
	ServiceName string
}

type OtelMiddleware struct {
	Config *OtelMiddlewareConfig
}

func NewOtelMiddleware(config *OtelMiddlewareConfig) *OtelMiddleware {
	return &OtelMiddleware{
		Config: config,
	}
}

func (instance *OtelMiddleware) Handle(c *gin.Context) {
	otelgin.Middleware(instance.Config.ServiceName)(c)
}
