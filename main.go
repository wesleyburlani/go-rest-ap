package main

import (
	http_port "test/web-service/ports/http"
)

func main() {
	http_port.StartServer(8080)
}
