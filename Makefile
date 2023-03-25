client:
	go run ./cmd/client/main.go

build:
	go build -o client ./cmd/client/main.go

test:
	go test ./...
