# The first stage: compile notification service
FROM golang:1.13 as builder
LABEL name="watcher" maintainer="o.kaya" version="0.1"
ENV APP_NAME watcher
WORKDIR /opt/${APP_NAME}
COPY ../.. .
RUN CGO_ENABLED=0 go build -o ./bin/watcher .

# The second stage: copy the notification binary to another container
FROM alpine:3.9
WORKDIR /opt/watcher
COPY --from=builder /opt/go-calendar/bin/watcher .
COPY --from=builder /opt/go-calendar/configs/config.yaml ./configs/
CMD ["./watcher", "-c", "./configs/config.yaml"]
