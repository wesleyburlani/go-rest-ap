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

// GetAlbum 	godoc
// @Summary 	returns a album by its id
// @Schemes 	http https
// @Description returns a album by its id
// @Tags 		albums
// @Produce 	json
// @Param		id		path	uint  true  "album id"
// @Success 	200	{object} models.Album
// @Failure		400	{object} models.Error
// @Failure		404	{object} models.Error
// @Failure		500	{object} models.Error
// @Router 		/albums/{id} [get]
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
