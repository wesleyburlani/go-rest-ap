package albums_controller

import (
	"net/http"

	"github.com/wesleyburlani/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

type PutAlbumUriParams struct {
	Id string `uri:"id" binding:"required"`
}

type PutAlbumBody struct {
	Title  string  `json:"title" binding:""`
	Artist string  `json:"artist" binding:""`
	Price  float64 `json:"price" binding:"numeric"`
}

func (instance *AlbumsController) PutAlbum(c *gin.Context) {
	uri := PutAlbumUriParams{}
	body := PutAlbumBody{}

	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	album, err := instance.albumsService.UpdateAlbum(uri.Id, models.AlbumProps{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	})

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, album)
}
