package http

import "github.com/gin-gonic/gin"

type Controller interface {
	Method() string
	RelativePath() string
	Handle(c *gin.Context)
}
