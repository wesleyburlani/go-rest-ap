package albums_controller

import (
	"net/http"

	"github.com/wesleyburlani/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

type PostAlbumBody struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required,gte=0"`
}

func (instance *AlbumsController) PostAlbum(c *gin.Context) {
	body := PostAlbumBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, instance.albumsService.CreateAlbum(models.AlbumProps{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	}))
}
