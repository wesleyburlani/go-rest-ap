package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	http_pkg "github.com/wesleyburlani/go-rest-api/pkg/http"
	null "gopkg.in/guregu/null.v4"
)

type Put struct {
	logger *logrus.Logger
	svc    *users.Service
}

type PutUri struct {
	ID int64 `uri:"id" binding:"required"`
}

type PutBody struct {
	Email    null.String `json:"email" swaggertype:"string"`
	Password null.String `json:"password,omitempty" swaggertype:"string"`
}

func NewPut(logger *logrus.Logger, svc *users.Service) *Put {
	return &Put{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *Put) Method() string {
	return "PUT"
}

func (ctl *Put) RelativePath() string {
	return "/users/:id"
}

func (ctl *Put) Middlewares() []http_pkg.Middleware {
	return []http_pkg.Middleware{}
}

// PutUser			godoc
// @Summary			updates an existing user
// @Schemes			http https
// @Description	updates an existing user
// @Tags				users
// @Produce			json
// @Param				id						path			int64				true	"user id"
// @Param				request				body			PutBody			true	"update user properties"
// @Success			200						{object}	users.User
// @Failure			400						{object}	string
// @Failure			500						{object}	string
// @Router			/users/{id} 	[put]
func (ctl *Put) Handle(ctx *gin.Context) {
	uri := PutUri{}
	body := PutBody{}

	if err := ctx.BindUri(&uri); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			ctl.logger.Debugf("error aborting request %v\n", err)
		}
		return
	}

	user, err := ctl.
		svc.
		WithContext(ctx.Request.Context()).
		Update(uri.ID, users.UpdateUserProps{
			Email: sql.NullString{
				String: body.Email.String,
				Valid:  body.Email.Valid,
			},
			Password: sql.NullString{
				String: body.Password.String,
				Valid:  body.Password.Valid,
			},
		})

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}
