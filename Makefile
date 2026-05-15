# Recursive Snake Case (rsnac) Makefile
# ==================================================

.PHONY: all fmt vet test build run clean

all: fmt vet test

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

build:
	go build -o rsnac ./cmd/rsnac/main.go

#run:
#	go run ./cmd/rsnac/main.go

clean:
	go clean
