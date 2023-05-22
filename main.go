package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

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

	var server *gin.Engine
	err = container.Resolve(&server)

	if err != nil {
		panic(err)
	}

	var config *config.Config
	err = container.Resolve(&config)
	if err != nil {
		panic(err)
	}

	server.Run(fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort))
}
