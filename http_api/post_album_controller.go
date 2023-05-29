package http_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/models"
	"github.com/wesleyburlani/go-rest-api/services"
)

type PostAlbumController struct {
	albumsService services.IAlbumsService
}

func NewPostAlbumController(
	albumsService services.IAlbumsService,
) *PostAlbumController {
	return &PostAlbumController{
		albumsService,
	}
}

type PostAlbumBody struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required,gte=0"`
}

func (instance *PostAlbumController) Method() string {
	return "POST"
}

func (instance *PostAlbumController) RelativePath() string {
	return "/albums"
}

func (instance *PostAlbumController) Handle(c *gin.Context) {
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
