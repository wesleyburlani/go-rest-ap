package users

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	"github.com/wesleyburlani/go-rest-api/internal/users"
)

type Post struct {
	logger *logrus.Logger
	svc    *users.Service
}

type PostBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

func NewPost(logger *logrus.Logger, svc *users.Service) *Post {
	return &Post{
		logger: logger,
		svc:    svc,
	}
}

func (ctl *Post) Method() string {
	return "POST"
}

func (ctl *Post) RelativePath() string {
	return "/users"
}

// PostUser			godoc
// @Summary			creates a new user
// @Schemes			http https
// @Description	creates a new user
// @Tags				users
// @Produce			json
// @Param				request	body			PostBody	true	"album properties"
// @Success			201			{object}	users.User
// @Failure			400			{object}	http.HTTPError
// @Failure			500			{object}	http.HTTPError
// @Router			/users 	[post]
func (ctl *Post) Handle(ctx *gin.Context) {
	body := PostBody{}

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
		Create(users.CreateUserProps{
			Email:    sql.NullString{String: body.Email, Valid: body.Email != ""},
			Password: sql.NullString{String: body.Password, Valid: body.Password != ""},
		})

	if err != nil {
		http_server.HandleError(ctx, err)
	} else {
		ctx.JSON(http.StatusCreated, user)
	}
}
