package http_port

import (
	"fmt"
	"test/web-service/database"
	albums_controller "test/web-service/ports/http/controllers/albums"
	albums_service "test/web-service/services/albums"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors)
	}
}

func StartServer(port int16) {
	router := gin.Default()
	router.Use(ErrorHandler)
	database := database.NewDatabase()
	albums_controller.NewAlbumsController(albums_service.NewAlbumsService(database)).RegisterRoutes(router)
	router.Run(fmt.Sprintf("%s:%d", "localhost", port))
}
