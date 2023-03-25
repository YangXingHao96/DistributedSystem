Combined golang repository for server, client and common marshal/unmarshal interfaces
### Client Usage

Build binary
```
make build
```
or
```
go build -o client ./cmd/client/main.go
```

Running client (Assuming server is already up)
#### Without flags
```
./client
```

#### Overriding default host and port of server
```
./client --host=localhost --port=3222
```
