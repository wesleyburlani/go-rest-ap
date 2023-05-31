package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/models"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

type PutAlbumController struct {
	albumsService service_albums.IAlbumsService
}

func NewPutAlbumController(
	albumsService service_albums.IAlbumsService,
) *PutAlbumController {
	return &PutAlbumController{
		albumsService,
	}
}

func (instance *PutAlbumController) Method() string {
	return "PUT"
}

func (instance *PutAlbumController) RelativePath() string {
	return "/albums/:id"
}

type PutAlbumUriParams struct {
	Id uint `uri:"id" binding:"required"`
}

type PutAlbumBody struct {
	Title  string  `json:"title" binding:""`
	Artist string  `json:"artist" binding:""`
	Price  float64 `json:"price" binding:"numeric"`
}

// PostAlbum 	godoc
// @Summary 	updates an album
// @Schemes 	http https
// @Description updates an album
// @Tags 		albums
// @Produce 	json
// @Param		id		path	uint  				true 	"album id"
// @Param 		request body 	models.AlbumProps 	true 	"album properties"
// @Success 	201	{object} models.Album
// @Failure		400	{object} models.Error
// @Failure		404	{object} models.Error
// @Failure		500	{object} models.Error
// @Router 		/albums/{id} [put]
func (instance *PutAlbumController) Handle(c *gin.Context) {
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
