package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/models"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

type PostAlbumController struct {
	logger        *logrus.Logger
	albumsService service_albums.IAlbumsService
}

func NewPostAlbumController(
	logger *logrus.Logger,
	albumsService service_albums.IAlbumsService,
) *PostAlbumController {
	return &PostAlbumController{
		logger,
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

// PostAlbum		godoc
// @Summary			creates a new album
// @Schemes			http https
// @Description	creates a new album
// @Tags				albums
// @Produce			json
// @Param				request	body			models.AlbumProps	true	"album properties"
// @Success			201			{object}	models.Album
// @Failure			400			{object}	models.Error
// @Failure			500			{object}	models.Error
// @Router			/albums [post]
func (instance *PostAlbumController) Handle(c *gin.Context) {
	body := PostAlbumBody{}

	if err := c.BindJSON(&body); err != nil {
		err := c.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			instance.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	c.JSON(http.StatusCreated, instance.
		albumsService.
		WithContext(c.Request.Context()).
		Create(models.AlbumProps{
			Title:  body.Title,
			Artist: body.Artist,
			Price:  body.Price,
		}))
}
