package di

import (
	http_port "github.com/wesleyburlani/go-rest-api/ports/http"
	http_controllers "github.com/wesleyburlani/go-rest-api/ports/http/controllers"
	albums_controller "github.com/wesleyburlani/go-rest-api/ports/http/controllers/albums"
	http_middlewares "github.com/wesleyburlani/go-rest-api/ports/http/middlewares"
	albums_service "github.com/wesleyburlani/go-rest-api/services/albums"

	logger_utils "github.com/wesleyburlani/go-rest-api/utils/logger/hooks"

	"github.com/wesleyburlani/go-rest-api/database"

	"github.com/goava/di"
	"github.com/sirupsen/logrus"
)

func BuildContainer() (*di.Container, error) {
	di.SetTracer(&di.StdTracer{})
	container, err := di.New(
		di.Provide(func() *logrus.Logger {
			logger := logrus.New()
			logger.Formatter = &logrus.JSONFormatter{}
			logger.AddHook(&logger_utils.TelemetryHook{})
			return logger
		}),
		di.Provide(http_port.NewHttpServer),
		// otel middleware must be the first on to be imported
		di.Provide(http_middlewares.NewOtelMiddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(http_middlewares.NewLoggerMiddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(http_middlewares.NewErrorMiddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(albums_controller.NewAlbumsController, di.As(new(http_controllers.IController))),
		di.Provide(albums_service.NewAlbumsService, di.As(new(albums_service.IAlbumsService))),
		di.Provide(database.NewDatabase),
	)
	return container, err
}
