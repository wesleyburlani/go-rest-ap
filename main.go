package main

import (
	http_port "test/web-service/ports/http"
	http_controllers "test/web-service/ports/http/controllers"
	albums_controller "test/web-service/ports/http/controllers/albums"
	http_middlewares "test/web-service/ports/http/middlewares"
	albums_service "test/web-service/services/albums"

	"fmt"
	"test/web-service/database"

	"github.com/gin-gonic/gin"
	"github.com/goava/di"
)

func main() {
	di.SetTracer(&di.StdTracer{})

	container, err := di.New(
		di.Provide(http_port.NewHttpServer),
		di.Provide(http_middlewares.NewErrorMiddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(albums_controller.NewAlbumsController, di.As(new(http_controllers.IController))),
		di.Provide(albums_service.NewAlbumsService, di.As(new(albums_service.IAlbumsService))),
		di.Provide(database.NewDatabase),
	)

	var server *gin.Engine
	err = container.Resolve(&server)
	if err != nil {
		panic(err)
	}

	server.Run(fmt.Sprintf("%s:%d", "localhost", 8080))
}
