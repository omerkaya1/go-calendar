# The first stage: compile go-calendar
FROM golang:1.13-alpine as dependency-builder
ENV APP_NAME go-calendar
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

# The second stage:
FROM dependency-builder as app-builder
ENV APP_NAME go-calendar
WORKDIR /opt/${APP_NAME}
COPY --from=dependency-builder /opt/go-calendar .
RUN CGO_ENABLED=0 go build -o ./bin/go-calendar .

# The third stage: copy the go-calendar binary to another container
FROM alpine:3.9
LABEL name="go-calendar" maintainer="o.kaya" version="0.1"
WORKDIR /opt/go-calendar
COPY --from=app-builder /opt/go-calendar/bin/go-calendar ./bin/
COPY --from=app-builder /opt/go-calendar/configs/config.json ./configs/
EXPOSE 7070
CMD ["./bin/go-calendar", "grpc-server", "-c", "./configs/config.json"]
