package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyburlani/go-rest-api/internal/users"
)

func GetUserFromContext(c *gin.Context) *users.User {
	user, ok := c.Get("user")
	if !ok {
		return nil
	}
	return user.(*users.User)
}
