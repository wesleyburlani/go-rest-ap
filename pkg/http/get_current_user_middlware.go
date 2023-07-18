package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
)

type GetCurrentUserMiddleware struct {
	svc *users.Service
}

func NewGetCurrentUserMiddleware(svc *users.Service) *GetCurrentUserMiddleware {
	return &GetCurrentUserMiddleware{
		svc: svc,
	}
}

func (m *GetCurrentUserMiddleware) Handle(c *gin.Context) {
	props, exists := c.Get("jwt")

	if !exists {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	jwtProps := props.(*crypto.JwtProps)
	user, err := m.svc.GetByEmail(jwtProps.Username)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Set("user", user)
	c.Next()
}
