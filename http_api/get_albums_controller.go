package http_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/services"
)

type GetAlbumsController struct {
	albumsService services.IAlbumsService
}

func NewGetAlbumsController(
	albumsService services.IAlbumsService,
) *GetAlbumsController {
	return &GetAlbumsController{
		albumsService,
	}
}

type GetAlbumsQueryParams struct {
	Page  int `form:"page,default=0" binding:"numeric,gte=0"`
	Limit int `form:"limit,default=20" binding:"numeric,gte=0"`
}

func (instance *GetAlbumsController) Method() string {
	return "GET"
}

func (instance *GetAlbumsController) RelativePath() string {
	return "/albums"
}

// GetAlbums 	godoc
// @Summary 	returns a list of albums
// @Schemes 	http https
// @Description returns a list of albums
// @Tags 		albums
// @Produce 	json
// @Param		page	query	int  false  "page"
// @Param		limit	query	int  false  "limit"
// @Success 	200		{object} []models.Album
// @Failure		400		{object} models.Error
// @Failure		500		{object} models.Error
// @Router 		/albums [get]
func (instance *GetAlbumsController) Handle(c *gin.Context) {
	params := GetAlbumsQueryParams{}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, instance.albumsService.GetAlbums(params.Page, params.Limit))
}
