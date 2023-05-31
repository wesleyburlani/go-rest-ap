package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/utils"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type OtelMiddleware struct {
	Config *utils.Config
}

func NewOtelMiddleware(cfx *utils.Config) *OtelMiddleware {
	return &OtelMiddleware{
		Config: cfx,
	}
}

func (instance *OtelMiddleware) Handle(c *gin.Context) {
	otelgin.Middleware(instance.Config.ServiceName)(c)
}
