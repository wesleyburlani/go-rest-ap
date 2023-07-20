package di

import (
	"github.com/goava/di"
	_ "github.com/lib/pq"
	auth_service "github.com/wesleyburlani/go-rest-api/internal/auth"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	http_server "github.com/wesleyburlani/go-rest-api/internal/transport/http"
	auth_controllers "github.com/wesleyburlani/go-rest-api/internal/transport/http/controllers/auth"
	users_controllers "github.com/wesleyburlani/go-rest-api/internal/transport/http/controllers/users"
	"github.com/wesleyburlani/go-rest-api/internal/transport/http/middlewares"
	users_service "github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
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

	auth := di.Options(
		di.Provide(func() *crypto.JwtAuth {
			return crypto.NewJwtAuth([]byte(cfg.JwtSecretKey))
		}),
	)

	httpServer := di.Options(
		// otel middleware must be the first one to be imported
		di.Provide(func() *middlewares.OtelMiddlewareConfig {
			return &middlewares.OtelMiddlewareConfig{
				ServiceName: cfg.ServiceName,
			}
		}),

		// middlewares
		di.Provide(middlewares.NewOtelMiddleware, di.As(new(http.Middleware))),

		// controllers
		di.Provide(auth_controllers.NewLogin, di.As(new(http.Controller))),

		di.Provide(users_controllers.NewPost, di.As(new(http.Controller))),
		di.Provide(users_controllers.NewPut, di.As(new(http.Controller))),
		di.Provide(users_controllers.NewGet, di.As(new(http.Controller))),
		di.Provide(users_controllers.NewDelete, di.As(new(http.Controller))),
		di.Provide(users_controllers.NewList, di.As(new(http.Controller))),
		di.Provide(users_controllers.NewMe, di.As(new(http.Controller))),

		di.Provide(http_server.NewServer, di.Tags{"type": "http"}),
	)

	websocketServer := di.Options()

	services := di.Options(
		di.Provide(db.NewDatabase),
		di.Provide(users_service.NewService, di.As(new(users_service.IService))),
		di.Provide(auth_service.NewService, di.As(new(auth_service.IService))),
	)

	container, err := di.New(
		utils,
		auth,
		httpServer,
		websocketServer,
		services,
	)

	return container, err
}
