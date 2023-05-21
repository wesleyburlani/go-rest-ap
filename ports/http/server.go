package http_port

import (
	http_controllers "test/web-service/ports/http/controllers"
	http_middlewares "test/web-service/ports/http/middlewares"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(middlewares []http_middlewares.IMiddleware, controllers []http_controllers.IController) *gin.Engine {
	router := gin.Default()
	for _, middleware := range middlewares {
		router.Use(middleware.Handle)
	}

	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
	return router
}
