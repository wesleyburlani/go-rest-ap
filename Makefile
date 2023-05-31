# starts to serve the api 
start:
	make -B swagger && go run .
# execute all tests on the repository
test:
	go test ./...
# updates swagger docs based on the latest code
swagger:
	swag init -g http_api/server.go --output swagger
# builds the application and outputs to bin/ folder
build:
	make -B swagger && go build -o bin/app
# deletes the contents of bin/ folder
clean:
	rm -Rf bin/*