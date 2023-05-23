package http_middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type OtelMiddleware struct {
	Config *config.Config
}

func NewOtelMiddleware(cfx *config.Config) *OtelMiddleware {
	return &OtelMiddleware{
		Config: cfx,
	}
}

func (instance *OtelMiddleware) Handle(c *gin.Context) {
	otelgin.Middleware(instance.Config.ServiceName)(c)
}
