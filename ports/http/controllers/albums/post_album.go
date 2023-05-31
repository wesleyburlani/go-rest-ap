package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/models"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

type PostAlbumController struct {
	albumsService service_albums.IAlbumsService
}

func NewPostAlbumController(
	albumsService service_albums.IAlbumsService,
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

// PostAlbum 	godoc
// @Summary 	creates a new album
// @Schemes 	http https
// @Description creates a new album
// @Tags 		albums
// @Produce 	json
// @Param 		request body 	models.AlbumProps 	true 	"album properties"
// @Success 	201	{object} models.Album
// @Failure		400	{object} models.Error
// @Failure		500	{object} models.Error
// @Router 		/albums [post]
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
