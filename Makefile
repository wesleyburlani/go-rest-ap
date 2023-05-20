# starts to serve the api 
start:
	go run .
# execute all tests on the repository
test:
	go test ./...
# builds the application and outputs to bin/ folder
build:
	go build -o bin/app
# deletes the contents of bin/ folder
clean:
	rm -Rf bin/*