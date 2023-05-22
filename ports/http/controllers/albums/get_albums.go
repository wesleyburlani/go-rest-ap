package albums_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAlbumsQueryParams struct {
	Page  int `form:"page,default=0" binding:"numeric,gte=0"`
	Limit int `form:"limit,default=20" binding:"numeric,gte=0"`
}

func (instance *AlbumsController) GetAlbums(c *gin.Context) {
	ctx := c.Request.Context()
	instance.logger.WithContext(ctx).Error("test")

	params := GetAlbumsQueryParams{}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, instance.albumsService.GetAlbums(params.Page, params.Limit))
}
