package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type Me struct {
	logger *logrus.Logger
	svc    *users.Service
	auth   *crypto.JwtAuth
}

func NewMe(logger *logrus.Logger, svc *users.Service, auth *crypto.JwtAuth) *Me {
	return &Me{
		logger: logger,
		svc:    svc,
		auth:   auth,
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
		http_pkg.NewJwtMiddleware(ctl.auth),
		http_pkg.NewGetCurrentUserMiddleware(ctl.svc),
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
	user := ctx.MustGet("user").(users.User)
	ctx.JSON(200, user)
}
