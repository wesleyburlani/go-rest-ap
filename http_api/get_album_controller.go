package http_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/services"
)

type GetAlbumController struct {
	albumsService services.IAlbumsService
}

func NewGetAlbumController(
	albumsService services.IAlbumsService,
) *GetAlbumController {
	return &GetAlbumController{
		albumsService,
	}
}

type GetAlbumUriParams struct {
	Id uint `uri:"id" binding:"required"`
}

func (instance *GetAlbumController) Method() string {
	return "GET"
}

func (instance *GetAlbumController) RelativePath() string {
	return "/albums/:id"
}

func (instance *GetAlbumController) Handle(c *gin.Context) {
	uri := GetAlbumUriParams{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	album, err := instance.albumsService.GetAlbum(uri.Id)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, album)
}
