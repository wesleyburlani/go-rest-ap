package main

import (
	"context"
	"log"
	http_port "test/web-service/ports/http"
	http_controllers "test/web-service/ports/http/controllers"
	albums_controller "test/web-service/ports/http/controllers/albums"
	http_middlewares "test/web-service/ports/http/middlewares"
	albums_service "test/web-service/services/albums"

	logger_utils "test/web-service/utils/logger/hooks"

	"fmt"
	"test/web-service/database"

	"github.com/gin-gonic/gin"
	"github.com/goava/di"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

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
		di.Provide(http_middlewares.NewLoggeriddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(http_middlewares.NewErrorMiddleware, di.As(new(http_middlewares.IMiddleware))),
		di.Provide(albums_controller.NewAlbumsController, di.As(new(http_controllers.IController))),
		di.Provide(albums_service.NewAlbumsService, di.As(new(albums_service.IAlbumsService))),
		di.Provide(database.NewDatabase),
	)

	if err != nil {
		panic(err)
	}

	var server *gin.Engine
	err = container.Resolve(&server)
	if err != nil {
		panic(err)
	}

	server.Run(fmt.Sprintf("%s:%d", "localhost", 8080))
}
