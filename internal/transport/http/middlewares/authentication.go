package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/internal/auth"
)

type AuthenticationMiddleware struct {
	authSvc auth.IService
}

func NewAuthenticationMiddleware(authSvc auth.IService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		authSvc: authSvc,
	}
}

func (m *AuthenticationMiddleware) Handle(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := ""
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	}

	user, err := m.authSvc.GetUserFromJwtToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("user", user)
	c.Next()
}
