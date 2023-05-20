package albums_controller_test

import (
	albums_controller "test/web-service/ports/http/controllers/albums"
	albums_service "test/web-service/services/albums"

	"github.com/gin-gonic/gin"
)

var mockAlbumsService = albums_service.NewMockAlbumsService()

func setupRouter() *gin.Engine {
	controller := albums_controller.NewAlbumsController(mockAlbumsService)
	router := gin.Default()
	controller.RegisterRoutes(router)
	return router
}
