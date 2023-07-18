package http

import "github.com/gin-gonic/gin"

type Controller interface {
	Method() string
	Middlewares() []Middleware
	RelativePath() string
	Handle(c *gin.Context)
}
