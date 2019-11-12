# The first stage: compile integration-test
FROM golang:1.13-alpine as dependency-builder
ENV APP_NAME integration-test
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

# The second stage:
FROM dependency-builder as app-builder
ENV APP_NAME integration-test
WORKDIR /opt/${APP_NAME}
COPY --from=dependency-builder /opt/integration-test .
RUN CGO_ENABLED=0 go build -o ./bin/integration-test ./cmd/integration-test

# The third stage: copy the integration-test binary to another container
FROM alpine:3.9
LABEL name="integration-test" maintainer="o.kaya" version="0.1"
WORKDIR /opt/integration-test
COPY --from=app-builder /opt/integration-test/bin/integration-test ./bin/
COPY --from=app-builder /opt/integration-test/configs/config.json ./configs/
COPY --from=app-builder /opt/integration-test/test/ ./test/
