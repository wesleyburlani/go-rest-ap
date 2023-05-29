package http_api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	for _, controller := range controllers {
		router.Handle(controller.Method(), controller.RelativePath(), controller.Handle)
	}
	return router
}
