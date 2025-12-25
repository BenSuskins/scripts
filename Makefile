.PHONY: all build run test fmt vet tidy clean

all: build

build:
	go build ./cmd/scripts

run:
	go run ./cmd/scripts

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -f scripts
