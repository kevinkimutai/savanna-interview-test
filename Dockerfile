# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22.3-alpine3.19 AS build-stage

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -cover ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main /main
COPY --from=build-stage /app/.env /.env
COPY --from=build-stage /app/docs /docs

EXPOSE 8080


ENTRYPOINT ["/main"]