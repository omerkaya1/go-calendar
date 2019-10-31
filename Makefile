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

.PHONY: mod
mod: ## Runs mod
	go mod verify
	go mod vendor
	go mod tidy

.PHONY: test
test: setup ## Runs all the tests
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Runs all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: setup ## Runs goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: setup ## Runs all the linters
	golint ./internal ./cmd ./configs ./log ./

.PHONY: build
build: ## Builds the project
	go build -o $(BUILD)/go-calendar $(CURDIR)

.PHONY: gen
gen: ## Triggers code generation of
	protoc --go_out=plugins=grpc:$(CURDIR)/internal/grpc api/*.proto

.PHONY: dockerbuild-gc
dockerbuild-gc: ## Builds a docker image with a project
	docker build -t omer513/go-calendar:0.${VERSION} -f ./deployments/go-calendar/Dockerfile .

.PHONY: dockerpush-gc
dockerpush-gc: dockerbuild-gc ## Publishes the docker image to the registry
	docker push omer513/go-calendar:0.${VERSION}

.PHONY: docker-compose-up
docker-compose-up: ## Runs docker-compose command to kick-start the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml up -d

.PHONY: docker-compose-down
docker-compose-down: ## Runs docker-compose command to remove the turn down the infrastructure
	docker-compose -f ./deployments/docker-compose.yaml down -v

.PHONY: clean
clean: ## Remove temporary files
	go clean $(CURDIR)
	rm -rf $(BUILD)
	rm -rf coverage.txt

.PHONY: help
help: ## Print this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
