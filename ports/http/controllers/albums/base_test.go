package albums_controller_test

import (
	albums_controller "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	albums_service "github.com/wesleyburlani/go-rest-api/services/albums"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var mockAlbumsService = albums_service.NewMockAlbumsService()

func setupRouter() *gin.Engine {
	controller := albums_controller.NewAlbumsController(logrus.New(), mockAlbumsService)
	router := gin.Default()
	controller.RegisterRoutes(router)
	return router
}
