package http_port

import (
	"github.com/wesleyburlani/go-rest-api/config"
	http_controllers "github.com/wesleyburlani/go-rest-api/ports/http/controllers"
	http_middlewares "github.com/wesleyburlani/go-rest-api/ports/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewHttpServer(
	middlewares []http_middlewares.IMiddleware,
	controllers []http_controllers.IController,
	logger *logrus.Logger,
	cfg *config.Config,
) *gin.Engine {

	if cfg.Mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

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
