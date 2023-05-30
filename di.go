package main

import (
	"github.com/wesleyburlani/go-rest-api/db"
	"github.com/wesleyburlani/go-rest-api/http_api"
	"github.com/wesleyburlani/go-rest-api/services"
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

		di.Provide(http_api.NewGetAlbumController, di.As(new(http_api.Controller))),
		di.Provide(http_api.NewGetAlbumsController, di.As(new(http_api.Controller))),
		di.Provide(http_api.NewPostAlbumController, di.As(new(http_api.Controller))),
		di.Provide(http_api.NewPutAlbumController, di.As(new(http_api.Controller))),
		di.Provide(http_api.NewServer),

		di.Provide(http_api.NewOtelMiddleware, di.As(new(http_api.Middleware))),
		di.Provide(http_api.NewBeforeRequestLoggerMiddleware, di.As(new(http_api.Middleware))),
		di.Provide(http_api.NewAfterRquestLoggerMiddleware, di.As(new(http_api.Middleware))),
		di.Provide(http_api.NewErrorMiddleware, di.As(new(http_api.Middleware))),

		di.Provide(services.NewAlbumsService, di.As(new(services.IAlbumsService))),
	)

	return container, err
}
