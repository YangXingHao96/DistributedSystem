database:
	cd setup/ && docker-compose up -d
databaseDown:
	cd setup/ && docker-compose down -v
client:
	go run ./cmd/client/main.go
serverAloNoTimeout:
	go run ./cmd/server/main.go
serverAloTimeout:
	go run ./cmd/server/main.go -timeout=true
serverAmoNoTimeout:
	go run ./cmd/server/main.go -mode=1
serverAmoTimeout:
	go run ./cmd/server/main.go -timeout=true -mode=1
buildClient:
	go build -o client ./cmd/client/main.go
buildServer:
	go build -o server ./cmd/server/main.go
test:
	go test ./...
