package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
)

type JwtMiddleware struct {
	jwt *crypto.JwtAuth
}

func NewJwtMiddleware(jwt *crypto.JwtAuth) *JwtMiddleware {
	return &JwtMiddleware{
		jwt: jwt,
	}
}

func (m *JwtMiddleware) Handle(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	token := ""
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	}

	props, err := m.jwt.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("jwt", props)
	c.Next()
}
