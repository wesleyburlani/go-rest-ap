package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/goava/di"
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

	err = container.Invoke(func(config *utils.Config, logger *logrus.Logger) {

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

		wg.Add(1)
		go func() {
			var webSocketServer *gin.Engine
			err := container.Resolve(&webSocketServer, di.Tags{"type": "websocket"})
			if err != nil {
				panic(err)
			}
			address := fmt.Sprintf("%s:%d", config.WebSocketHost, config.WebSocketPort)
			logger.WithFields(logrus.Fields{
				"address": address,
			}).Info("websocket server running")
			if err = webSocketServer.Run(address); err != nil {
				panic(err)
			}
		}()

		wg.Wait()
	})

	if err != nil {
		panic(err)
	}
}
