ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: lint vet test coverage

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out

.PHONY: integration
integration:
	cd atdd && godog -f cucumber >&1 | tee $(ROOT_DIR)/cucumber.json

.PHONY: coverage
coverage: 
	go tool cover -html=coverage.out ; go tool cover -func=coverage.out

.PHONY: lint
lint:
	golint $(shell find * -type d | egrep -v '^vendor')

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: start
# TODO I could create a thin local RESTful server that proxys the Lambda's RPC as may benefit workflows that include POSTMAN, VSCode's REST Client, cURL, etc.
start:
	_LAMBDA_SERVER_PORT=8001 go run cmd/handler/handler.go

.PHONY: cleanup
cleanup:
	rm handler || true
	rm cucumber.json || true
	rm godog_dependency_file_test.go || true
	rm coverage.out || true
	rm -r vendor || true

.PHONY: build
build: 
	go build -o .bin/handler cmd/handler.go