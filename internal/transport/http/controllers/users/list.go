package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type List struct {
	logger *logrus.Logger
	svc    *users.Service
}

type ListQuery struct {
	Page  int32 `form:"page,default=0" binding:"numeric,gte=0"`
	Limit int32 `form:"limit,default=20" binding:"numeric,gte=0,lte=100"`
}

func NewList(logger *logrus.Logger, svc *users.Service) *List {
	return &List{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *List) Method() string {
	return "GET"
}

func (ctl *List) RelativePath() string {
	return "/users"
}

func (ctl *List) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{}
}

// ListUser			godoc
// @Summary			lists existing users
// @Schemes			http https
// @Description	lists existing users
// @Tags				users
// @Produce			json
// @Success			200						{object}	[]users.User
// @Failure			500						{object}	string
// @Router			/users 				[get]
func (ctl *List) Handle(ctx *gin.Context) {
	query := ListQuery{}

	if err := ctx.BindQuery(&query); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	user, err := ctl.
		svc.
		WithContext(ctx.Request.Context()).
		List(query.Page, query.Limit)

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}
