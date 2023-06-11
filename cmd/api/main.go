package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/goava/di"

	"github.com/sirupsen/logrus"

	"github.com/wesleyburlani/go-rest-api/internal/config"
	container_di "github.com/wesleyburlani/go-rest-api/internal/di"
	"github.com/wesleyburlani/go-rest-api/pkg/otel"
)

func main() {
	cfg := config.LoadConfig()

	tracerProvider := otel.Init(&otel.OtelConfig{
		ServiceName: cfg.ServiceName,
	})
	defer otel.Stop(tracerProvider)

	container, err := container_di.BuildContainer(cfg)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(config *config.Config, logger *logrus.Logger) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			var httpServer *gin.Engine
			err := container.Resolve(&httpServer, di.Tags{"type": "http"})
			if err != nil {
				panic(err)
			}
			address := fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort)
			logger.WithFields(logrus.Fields{
				"address": address,
			}).Info("http server running")
			if err = httpServer.Run(address); err != nil {
				panic(err)
			}
		}()

		wg.Wait()
	})

	if err != nil {
		panic(err)
	}
}
