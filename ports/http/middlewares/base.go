package http_middlewares

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	Handle(c *gin.Context)
}
