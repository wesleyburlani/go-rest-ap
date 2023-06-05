package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
)

type ListAlbumsController struct {
	logger        *logrus.Logger
	albumsService service_albums.IAlbumsService
}

func NewListAlbumsController(
	logger *logrus.Logger,
	albumsService service_albums.IAlbumsService,
) *ListAlbumsController {
	return &ListAlbumsController{
		logger,
		albumsService,
	}
}

type ListAlbumsQueryParams struct {
	Page  int `form:"page,default=0" binding:"numeric,gte=0"`
	Limit int `form:"limit,default=20" binding:"numeric,gte=0"`
}

func (instance *ListAlbumsController) Method() string {
	return "GET"
}

func (instance *ListAlbumsController) RelativePath() string {
	return "/albums"
}

// GetAlbums		godoc
// @Summary			returns a list of albums
// @Schemes			http https
// @Description	returns a list of albums
// @Tags				albums
// @Produce			json
// @Param				page		query			int							false	"page"
// @Param				limit		query			int							false	"limit"
// @Success			200			{object}	[]models.Album
// @Failure			400			{object}	models.Error
// @Failure			500			{object}	models.Error
// @Router			/albums	[get]
func (instance *ListAlbumsController) Handle(c *gin.Context) {
	params := ListAlbumsQueryParams{}

	if err := c.BindQuery(&params); err != nil {
		err := c.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			instance.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	c.JSON(
		http.StatusOK,
		instance.
			albumsService.
			WithContext(c.Request.Context()).
			List(params.Page, params.Limit),
	)
}
