# Build the application from source
FROM golang:1.20 as build

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app

# Run the tests in the container
FROM build AS test

RUN make test

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

COPY --from=build /app/bin/app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]