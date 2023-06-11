package di

import (
	"github.com/goava/di"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/database"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	users_controllers "github.com/wesleyburlani/go-rest-api/internal/transport/http/controllers/users"
	users_service "github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/http"
	"github.com/wesleyburlani/go-rest-api/pkg/logger"
)

func BuildContainer(cfg *config.Config) (*di.Container, error) {
	if cfg.Mode != config.ReleaseMode {
		di.SetTracer(&di.StdTracer{})
	}

	utils := di.Options(
		di.Provide(func() *config.Config { return cfg }),
		di.Provide(func() *logger.LoggerConfig {
			return &logger.LoggerConfig{
				LogLevel: cfg.LogLevel,
			}
		},
		),
		di.Provide(logger.NewLogger),
	)

	httpServer := di.Options(
		// otel middleware must be the first one to be imported
		di.Provide(func() *http.OtelMiddlewareConfig {
			return &http.OtelMiddlewareConfig{
				ServiceName: cfg.ServiceName,
			}
		}),
		di.Provide(http.NewOtelMiddleware, di.As(new(http.Middleware))),

		di.Provide(users_controllers.NewPost, di.As(new(http.Controller))),

		di.Provide(http_server.NewServer, di.Tags{"type": "http"}),
	)

	websocketServer := di.Options()

	services := di.Options(
		di.Provide(database.Init),
		di.Provide(users_service.NewService),
	)

	container, err := di.New(
		utils,
		httpServer,
		websocketServer,
		services,
	)

	return container, err
}
