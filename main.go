package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/wesleyburlani/go-rest-api/config"
	di "github.com/wesleyburlani/go-rest-api/di"
)

func main() {
	tracerProvider := initTracer()
	defer stopTracer(tracerProvider)

	container, err := di.BuildContainer()
	if err != nil {
		panic(err)
	}

	container.Invoke(func(server *gin.Engine, config *config.Config, logger *logrus.Logger) {
		address := fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort)
		logger.WithFields(logrus.Fields{
			"address": address,
		}).Info("server running")
		if err = server.Run(address); err != nil {
			panic(err)
		}
	})
}
