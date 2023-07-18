package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type Get struct {
	logger *logrus.Logger
	svc    *users.Service
}

type GetUri struct {
	ID int64 `uri:"id" binding:"required"`
}

func NewGet(logger *logrus.Logger, svc *users.Service) *Get {
	return &Get{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *Get) Method() string {
	return "GET"
}

func (ctl *Get) RelativePath() string {
	return "/users/:id"
}

func (ctl *Get) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{}
}

// GetUser			godoc
// @Summary			returns an existing user
// @Schemes			http https
// @Description	returns an existing user
// @Tags				users
// @Produce			json
// @Param				id						path			int64				true	"user id"
// @Success			200						{object}	users.User
// @Failure			404						{object}	string
// @Failure			500						{object}	string
// @Router			/users/{id} 	[get]
func (ctl *Get) Handle(ctx *gin.Context) {
	uri := GetUri{}

	if err := ctx.BindUri(&uri); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	user, err := ctl.
		svc.
		WithContext(ctx.Request.Context()).
		GetById(uri.ID)

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}
