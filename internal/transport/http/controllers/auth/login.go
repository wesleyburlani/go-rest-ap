package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/auth"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
)

type Login struct {
	logger *logrus.Logger
	svc    *auth.Service
}

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLogin(logger *logrus.Logger, svc *auth.Service) *Login {
	return &Login{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *Login) Method() string {
	return "POST"
}

func (ctl *Login) RelativePath() string {
	return "/auth/login"
}

func (ctl *Login) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{}
}

// PostUser			godoc
// @Summary			generates a jwt token for a user
// @Schemes			http https
// @Description	generates a jwt token for a user
// @Tags				auth
// @Produce			json
// @Param				request				body			PostBody		true	"login user properties"
// @Success			200						{object}	crypto.JwtToken
// @Failure			400						{object}	string
// @Failure			500						{object}	string
// @Router			/auth/login 	[post]
func (ctl *Login) Handle(ctx *gin.Context) {
	body := LoginBody{}

	if err := ctx.BindJSON(&body); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	var token *crypto.JwtToken
	token, err := ctl.
		svc.
		WithContext(ctx.Request.Context()).
		Login(body.Email, body.Password)

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, *token)
	}
}
