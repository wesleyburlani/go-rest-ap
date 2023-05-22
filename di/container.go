package di

import (
	http_port "test/web-service/ports/http"
	http_controllers "test/web-service/ports/http/controllers"
	albums_controller "test/web-service/ports/http/controllers/albums"
	http_middlewares "test/web-service/ports/http/middlewares"
	albums_service "test/web-service/services/albums"

	logger_utils "test/web-service/utils/logger/hooks"

	"test/web-service/database"

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
