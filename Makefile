BUILD= $(CURDIR)/bin
VERSION= $(shell git rev-list HEAD --count)
$(shell mkdir -p $(BUILD))
export GO111MODULE=on
export GOPATH=$(go env GOPATH)

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u golang.org/x/tools
	go get -u golang.org/x/lint/golint
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/DATA-DOG/godog/cmd/godog

.PHONY: mod
mod: ## Runs mod
	go mod verify
	go mod vendor
	go mod tidy

.PHONY: test
test: setup ## Runs all the tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: integration-test
integration-test: setup ## Runs integration tests
	go test -v -timeout=30s ./internal/go-calendar/grpc ./internal/mq/

.PHONY: cover
cover: test ## Runs all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: setup ## Runs goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: setup ## Runs all the linters
	golint ./internal ./cmd ./configs ./log ./

.PHONY: gen
gen: ## Triggers the protobuf code generation
	protoc --go_out=plugins=grpc:$(CURDIR)/internal/go-calendar/grpc api/*.proto

.PHONY: build-gcs
build-gcs: gen ## Builds the go-calendar project
	go build -o $(BUILD)/go-calendar $(CURDIR)/cmd/go-calendar

.PHONY: build-notification
build-notification: ## Builds the notification project
	go build -o $(BUILD)/notification $(CURDIR)/cmd/notification

.PHONY: build-watcher
build-watcher: ## Builds the watcher project
	go build -o $(BUILD)/watcher $(CURDIR)/cmd/watcher

.PHONY: build-all
build-all: build-gcs build-notification build-watcher ## Builds all binaries of the project

.PHONY: dockerbuild-gcs
dockerbuild-gcs: ## Builds a docker image with the go-calendar project
	docker build -t omer513/go-calendar:0.${VERSION} -f ./deployments/go-calendar/Dockerfile .

.PHONY: dockerpush-gcs
dockerpush-gcs: dockerbuild-gcs ## Publishes the docker image to the registry
	docker push omer513/go-calendar:0.${VERSION}

.PHONY: dockerbuild-notification
dockerbuild-notification: ## Builds a docker image with the notification project
	docker build -t omer513/notification:0.${VERSION} -f ./deployments/notification/Dockerfile .

.PHONY: dockerpush-notification
dockerpush-notification: dockerbuild-notification ## Publishes the docker image to the registry
	docker push omer513/notification:0.${VERSION}

.PHONY: dockerbuild-watcher
dockerbuild-watcher: ## Builds a docker image with a project
	docker build -t omer513/watcher:0.${VERSION} -f ./deployments/watcher/Dockerfile .

.PHONY: dockerpush-watcher
dockerpush-watcher: dockerbuild-watcher ## Publishes the docker image to the registry
	docker push omer513/watcher:0.${VERSION}

.PHONY: docker-compose-up
docker-compose-up: ## Runs docker-compose command to kick-start the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml up -d

.PHONY: docker-compose-down
docker-compose-down: ## Runs docker-compose command to remove the turn down the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml down -v

.PHONY: integration
integration: ##
	docker-compose -f ./deployments/docker-compose.test.yaml up --build -d;\
	test_status_code=0 ;\
	docker-compose -f ./deployments/docker-compose.test.yaml run integration_tests ./bin/integration-test || test_status_code=$$? ;\
	docker-compose -f ./deployments/docker-compose.test.yaml down --volumes;\
	printf "Return code is $$test_status_code\n" ;\
	exit $$test_status_code ;\

.PHONY: clean
clean: ## Remove temporary files
	go clean $(CURDIR)/cmd/go-calendar
	go clean $(CURDIR)/cmd/notification
	go clean $(CURDIR)/cmd/watcher
	go clean $(CURDIR)/cmd/integration-test
	rm -rf $(BUILD)
	rm -rf coverage.txt

.PHONY: help
help: ## Print this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build-gcs
