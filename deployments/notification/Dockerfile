# The first stage: compile notification
FROM golang:1.13-alpine as meta-builder
ENV APP_NAME notification
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

# The second stage:
FROM meta-builder as app-builder
ENV APP_NAME notification
WORKDIR /opt/${APP_NAME}
COPY --from=meta-builder /opt/notification .
RUN CGO_ENABLED=0 go build -o ./bin/notification ./cmd/notification

# The third stage: copy the notification binary to another container
FROM alpine:3.9
LABEL name="notification" maintainer="o.kaya" version="0.1"
WORKDIR /opt/notification
COPY --from=app-builder /opt/notification/bin/notification ./bin/
COPY --from=app-builder /opt/notification/configs/config.json ./configs/
CMD ["./bin/notification", "-c", "./configs/config.json"]
