package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/auth"
	"github.com/wesleyburlani/go-rest-api/internal/transport/http/middlewares"
	"github.com/wesleyburlani/go-rest-api/internal/transport/http/utils"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type Me struct {
	logger   *logrus.Logger
	usersSvc users.IService
	authSvc  auth.IService
}

func NewMe(
	logger *logrus.Logger,
	usersSvc users.IService,
	authSvc auth.IService,
) *Me {
	return &Me{
		logger:   logger,
		usersSvc: usersSvc,
		authSvc:  authSvc,
	}
}

func (ctl *Me) Method() string {
	return "GET"
}

func (ctl *Me) RelativePath() string {
	return "/users/me"
}

func (ctl *Me) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{
		middlewares.NewAuthenticationMiddleware(ctl.authSvc),
	}
}

// Me						godoc
// @Summary			returns the current authenticated user
// @Schemes			http https
// @Description	returns the current authenticated user
// @Tags				users
// @Produce			json
// @Success			200						{object}	users.User
// @Failure			404						{object}	string
// @Failure			500						{object}	string
// @Router			/users/me 		[get]
func (ctl *Me) Handle(ctx *gin.Context) {
	user := utils.GetUserFromContext(ctx)
	ctx.JSON(200, user)
}
