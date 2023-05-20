package albums_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAlbumUriParams struct {
	Id string `uri:"id" binding:"required"`
}

func (instance *AlbumsController) GetAlbum(c *gin.Context) {
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
	return
}
