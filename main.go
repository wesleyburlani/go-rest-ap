package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	di "test/web-service/di"
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

	server.Run(fmt.Sprintf("%s:%d", "localhost", 8080))
}
