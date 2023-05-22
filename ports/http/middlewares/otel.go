package http_middlewares

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type OtelMiddleware struct {
}

func NewOtelMiddleware() *OtelMiddleware {
	return &OtelMiddleware{}
}

func (instance *OtelMiddleware) Handle(c *gin.Context) {
	otelgin.Middleware("service-name")(c)
}
