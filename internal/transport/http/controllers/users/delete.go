package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type Delete struct {
	logger *logrus.Logger
	svc    *users.Service
}

type DeleteUri struct {
	ID int64 `uri:"id" binding:"required"`
}

func NewDelete(logger *logrus.Logger, svc *users.Service) *Delete {
	return &Delete{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *Delete) Method() string {
	return "DELETE"
}

func (ctl *Delete) RelativePath() string {
	return "/users/:id"
}

func (ctl *Delete) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{}
}

// DeleteUser		godoc
// @Summary			deletes an existing user
// @Schemes			http https
// @Description	deletes an existing user
// @Tags				users
// @Produce			json
// @Param				id						path			int64				true	"user id"
// @Success			200						{object}	string
// @Failure			404						{object}	string
// @Failure			500						{object}	string
// @Router			/users/{id} 	[Delete]
func (ctl *Delete) Handle(ctx *gin.Context) {
	uri := DeleteUri{}

	if err := ctx.BindUri(&uri); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	err := ctl.
		svc.
		WithContext(ctx.Request.Context()).
		Delete(uri.ID)

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, "")
	}
}
