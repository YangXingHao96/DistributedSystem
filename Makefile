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
server_amo_timeout:
	go run ./cmd/server/main.go -timeout=true -mode=1
build:
	go build -o client ./cmd/client/main.go
test:
	go test ./...
