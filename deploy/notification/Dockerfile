# The first stage: compile notification service
FROM golang:1.13 as builder
LABEL name="notification" maintainer="o.kaya" version="0.1"
ENV APP_NAME notification
WORKDIR /opt/${APP_NAME}
COPY ../.. .
RUN CGO_ENABLED=0 go build -o ./bin/notification .

# The second stage: copy the notification binary to another container
FROM alpine:3.9
WORKDIR /opt/notification
COPY --from=builder /opt/go-calendar/bin/notification .
COPY --from=builder /opt/go-calendar/configs/config.yaml ./configs/
CMD ["./notification", "-c", "./configs/config.yaml"]
