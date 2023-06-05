package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

type GetAlbumController struct {
	logger        *logrus.Logger
	albumsService service_albums.IAlbumsService
}

func NewGetAlbumController(
	logger *logrus.Logger,
	albumsService service_albums.IAlbumsService,
) *GetAlbumController {
	return &GetAlbumController{
		logger,
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

// GetAlbum			godoc
// @Summary			returns a album by its id
// @Schemes			http https
// @Description	returns a album by its id
// @Tags				albums
// @Produce			json
// @Param				id						path			uint					true	"album id"
// @Success			200						{object}	models.Album
// @Failure			400						{object}	models.Error
// @Failure			404						{object}	models.Error
// @Failure			500						{object}	models.Error
// @Router			/albums/{id} 	[get]
func (instance *GetAlbumController) Handle(c *gin.Context) {
	uri := GetAlbumUriParams{}

	if err := c.BindUri(&uri); err != nil {
		err := c.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			instance.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	album, err := instance.
		albumsService.
		WithContext(c.Request.Context()).
		Get(uri.Id)

	if err != nil {
		err := c.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			instance.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	c.JSON(http.StatusOK, album)
}
