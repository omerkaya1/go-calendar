# One-stage image building
FROM golang:1.13-alpine
LABEL name="go-calendar" maintainer="o.kaya" version="0.1"
WORKDIR /opt/go-calendar
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 go test -v -timeout=30s ./internal/go-calendar/grpc ./internal/mq/
