package http_port_controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes(r *gin.Engine)
}
