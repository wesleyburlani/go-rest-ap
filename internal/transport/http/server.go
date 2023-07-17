package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
	http_server "github.com/wesleyburlani/go-rest-api/pkg/http"
	docs "github.com/wesleyburlani/go-rest-api/swagger"
)

func abortWithStatusJson(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, http_server.HTTPError{
		Code:    code,
		Message: message,
	})
}

func HandleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	if instance, ok := err.(custom_errors.HttpError); ok {
		abortWithStatusJson(ctx, instance.StatusCode(), instance.Error())
		return
	}
	abortWithStatusJson(ctx, 500, err.Error())
}

func NewServer(
	middlewares []http_server.Middleware,
	controllers []http_server.Controller,
	cfg *config.Config,
) *gin.Engine {

	if cfg.Mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	for _, middleware := range middlewares {
		router.Use(middleware.Handle)
	}

	basePath := "/api/v1"

	v1 := router.Group(basePath)
	for _, controller := range controllers {
		v1.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	}

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
