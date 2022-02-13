.PHONY: all

all: clean build test

compile:
	go build -o ./bin ./main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint ./...

build: compile vet lint

clean:
	go clean

test:
	go test -v ./... -tags mock

generate-mocks:
	go get -d github.com/vektra/mockery/v2@v2.8.0
	mockery --all -r --inpackage 