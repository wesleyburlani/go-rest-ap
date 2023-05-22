package albums_controller

import (
	"fmt"
	albums_service "test/web-service/services/albums"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AlbumsController struct {
	logger        *logrus.Logger
	albumsService albums_service.IAlbumsService
}

func NewAlbumsController(logger *logrus.Logger, albumsService albums_service.IAlbumsService) *AlbumsController {
	return &AlbumsController{
		logger,
		albumsService,
	}
}

func (instance *AlbumsController) RegisterRoutes(router *gin.Engine) {
	routePrefix := "/albums"
	router.GET(fmt.Sprintf("%s%s", routePrefix, ""), instance.GetAlbums)
	router.GET(fmt.Sprintf("%s%s", routePrefix, "/:id"), instance.GetAlbum)
	router.POST(fmt.Sprintf("%s%s", routePrefix, ""), instance.PostAlbum)
	router.PUT(fmt.Sprintf("%s%s", routePrefix, "/:id"), instance.PutAlbum)
}
