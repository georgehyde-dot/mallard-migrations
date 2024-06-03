build:
	@go build -o ./cmd/main/main ./cmd/main

run: build
	@./cmd/main/main

test:
	@go test -v ./...

clean:
	@go fmt ./...
	@go mod tidy

.PHONY: clean test build run