package main

import (
	"github.com/wesleyburlani/go-rest-api/db"
	"github.com/wesleyburlani/go-rest-api/ports/http"
	http_controller_albums "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	http_middlewares "github.com/wesleyburlani/go-rest-api/ports/http/middlewares"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
	"github.com/wesleyburlani/go-rest-api/utils"

	"github.com/goava/di"
)

func BuildContainerDI() (*di.Container, error) {
	cfg := utils.LoadConfig()

	if cfg.Mode != utils.ReleaseMode {
		di.SetTracer(&di.StdTracer{})
	}
	container, err := di.New(
		di.Provide(func() *utils.Config { return cfg }),
		// ------- utils session -------
		di.Provide(utils.NewLogger),
		// otel middleware must be the first on to be imported

		di.Provide(db.Init),

		di.Provide(http_controller_albums.NewGetAlbumController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewGetAlbumsController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewPostAlbumController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewPutAlbumController, di.As(new(http.Controller))),
		di.Provide(http.NewServer),

		di.Provide(http_middlewares.NewOtelMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewBeforeRequestLoggerMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewAfterRquestLoggerMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewErrorMiddleware, di.As(new(http.Middleware))),

		di.Provide(service_albums.NewAlbumsService, di.As(new(service_albums.IAlbumsService))),
	)

	return container, err
}
