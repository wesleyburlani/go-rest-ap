package main

import (
	"github.com/wesleyburlani/go-rest-api/db"
	"github.com/wesleyburlani/go-rest-api/ports/http"
	http_controller_albums "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	http_middlewares "github.com/wesleyburlani/go-rest-api/ports/http/middlewares"
	"github.com/wesleyburlani/go-rest-api/ports/websocket"
	websocket_controller_echo "github.com/wesleyburlani/go-rest-api/ports/websocket/controllers/echo"
	websocket_middlewares "github.com/wesleyburlani/go-rest-api/ports/websocket/middlewares"
	service_albums "github.com/wesleyburlani/go-rest-api/services/albums"
	"github.com/wesleyburlani/go-rest-api/utils"

	"github.com/goava/di"
)

func BuildContainerDI(cfg *utils.Config) (*di.Container, error) {
	if cfg.Mode != utils.ReleaseMode {
		di.SetTracer(&di.StdTracer{})
	}

	utils := di.Options(
		di.Provide(func() *utils.Config { return cfg }),
		di.Provide(utils.NewLogger),
	)

	httpServer := di.Options(
		// otel middleware must be the first one to be imported
		di.Provide(http_middlewares.NewOtelMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewBeforeRequestLoggerMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewAfterRquestLoggerMiddleware, di.As(new(http.Middleware))),
		di.Provide(http_middlewares.NewErrorMiddleware, di.As(new(http.Middleware))),

		di.Provide(http_controller_albums.NewGetAlbumController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewGetAlbumsController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewPostAlbumController, di.As(new(http.Controller))),
		di.Provide(http_controller_albums.NewPutAlbumController, di.As(new(http.Controller))),

		di.Provide(http.NewServer, di.Tags{"type": "http"}),
	)

	websocketServer := di.Options(
		// otel middleware must be the first one to be imported
		di.Provide(websocket_middlewares.NewOtelMiddleware, di.As(new(websocket.Middleware))),
		di.Provide(websocket_middlewares.NewOnConnectionLoggerMiddleware, di.As(new(websocket.Middleware))),
		di.Provide(websocket_middlewares.NewOnDisconnectionLoggerMiddleware, di.As(new(websocket.Middleware))),

		di.Provide(websocket_controller_echo.NewWsEchoController, di.As(new(websocket.Controller))),

		di.Provide(websocket.NewServer, di.Tags{"type": "websocket"}),
	)

	services := di.Options(
		di.Provide(db.Init),
		di.Provide(service_albums.NewAlbumsService, di.As(new(service_albums.IAlbumsService))),
	)

	container, err := di.New(
		utils,
		httpServer,
		websocketServer,
		services,
	)

	return container, err
}
