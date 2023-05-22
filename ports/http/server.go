package http_port

import (
	http_controllers "github.com/wesleyburlani/go-rest-api/ports/http/controllers"
	http_middlewares "github.com/wesleyburlani/go-rest-api/ports/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewHttpServer(middlewares []http_middlewares.IMiddleware, controllers []http_controllers.IController, logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	//gin.DefaultWriter = logger.Writer()
	for _, middleware := range middlewares {
		router.Use(middleware.Handle)
	}

	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
	return router
}
