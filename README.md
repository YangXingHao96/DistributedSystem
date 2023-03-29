Combined golang repository for server, client and common marshal/unmarshal interfaces
### Client Usage

Build binary
```
make buildClient
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

### Server Usage

#### Set up database

up test database that will be used by the server
```
make database
```

down test database
```
make databaseDown
```
#### Running directly from DistributedSystem Directory, using default address localhost:2222

Running server with at least once delivery, without timeout simulation
```
make serverAloNoTimeout
```

Running server with at least once delivery, with timeout simulation
```
make serverAloTimeout
```

Running server with at most once delivery, without timeout simulation
```
make serverAmoNoTimeout
```

Running server with at most once delivery, without timeout simulation
```
make serverAmoNoTimeout
```

#### Build Binary

Build using makefile
```
make buildServer
```
Running server
#### Flags
mode -> integer 0 or 1, 0 implies server runs with at least once delivery, 1 implies server runs with at most once delivery, when no flag is set, default to 0

timeout ->boolean true or false, true implies server runs with timeout simulation, false implies server runs without timeout simulation, when no flag is set, defaults to false

host -> string, host address the server will listen to, default localhost

port -> string, port number that the server will listen to, default 2222

timeout% -> int, percentage that a server will simulate a timeout, default 20

#### Without flags
```
./server
```

#### With flags
Example:

run server in at most once delivery mode, timeout simulation, host address is localhost, listening on port 2222, 80 percent chance of timeout
```
./server -mode=1 -timeout=true -host=localhost -port=2000 -timeout%=80
```

### Running the client and server

order of creating: up database -> up server -> up clients

order of closing: close clients -> close server -> down database

