build:
	@go build -o bin/realtime-chat-backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/realtime-chat-backend