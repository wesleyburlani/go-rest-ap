# starts to serve the api
start:
	make -B swagger && go run cmd/api/main.go
# execute all tests on the repository
test:
	ENV=test go test -v ./...
# updates swagger docs based on the latest code
swagger:
	swag init -g internal/transport/http/server.go --output swagger
db-client:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate
# builds the application and outputs to bin/ folder
build:
	make -B db-client && make -B swagger && go build -o bin/api cmd/api/main.go
# deletes the contents of bin/ folder
clean:
	rm -Rf bin/*
