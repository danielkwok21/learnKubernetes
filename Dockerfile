# Multi-stage Dockerfile: build Go code from the build context (assumes build context is ./src)
FROM golang:1.20 AS builder

WORKDIR /src
COPY . ./

# build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /demo ./

EXPOSE 80
ENTRYPOINT ["/demo"]
