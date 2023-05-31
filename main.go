package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/wesleyburlani/go-rest-api/utils"
)

func main() {
	cfg := utils.LoadConfig()

	tracerProvider := initTracer(cfg)
	defer stopTracer(tracerProvider)

	container, err := BuildContainerDI(cfg)
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(server *gin.Engine, config *utils.Config, logger *logrus.Logger) {
		address := fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort)
		logger.WithFields(logrus.Fields{
			"address": address,
		}).Info("server running")
		if err = server.Run(address); err != nil {
			panic(err)
		}
	})

	if err != nil {
		panic(err)
	}
}
