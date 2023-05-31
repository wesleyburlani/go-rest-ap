package http_api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/wesleyburlani/go-rest-api/swagger"
	"github.com/wesleyburlani/go-rest-api/utils"
)

type Controller interface {
	Method() string
	RelativePath() string
	Handle(c *gin.Context)
}

type Middleware interface {
	Handle(c *gin.Context)
}

func NewServer(
	middlewares []Middleware,
	controllers []Controller,
	logger *logrus.Logger,
	cfg *utils.Config,
) *gin.Engine {

	if cfg.Mode == utils.ReleaseMode {
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
