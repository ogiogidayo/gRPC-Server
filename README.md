# gRPC-Server
This project is a simple implementation of a gRPC server in Go. It uses the `GreetingService` service which has four methods: `Hello`, `HelloServerStream`, `HelloClientStream`, and `HelloBiStreams`.

## Directory Structure
```
.
├── Makefile
├── README.md
├── api
│   └── hello.proto
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── go.mod
├── go.sum
└── pkg
    └── grpc
        ├── hello.pb.go
        └── hello_grpc.pb.go

```

## Services

### Hello (Unary RPC)

This is a simple RPC that accepts a `HelloRequest` and returns a `HelloResponse`. In Unary RPC, the client sends a single request to the server and gets a single response back, just like a normal function call.

### HelloServerStream (Server streaming RPC)

This is a server streaming RPC that accepts a `HelloRequest` and returns a stream of `HelloResponse`. In Server streaming RPC, the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages.

### HelloClientStream (Client streaming RPC)

This is a client streaming RPC that accepts a stream of `HelloRequest` and returns a `HelloResponse`. In Client streaming RPC, the client writes a sequence of messages and sends them to the server via a provided stream. Once the client has finished writing the messages, it waits for the server to read them all and return a response.

### HelloBiStreams (Bidirectional streaming RPC)

This is a bidirectional streaming RPC that accepts a stream of `HelloRequest` and returns a stream of `HelloResponse`. In Bidirectional streaming RPC, both the client and the server send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client messages before writing its responses, or it could alternately read a message then write a message, or some other combination of reads and writes.

## Setup

To run the server, execute the following command:

```bash
go run cmd/server/main.go
```
To run the client, execute the following command:
```bash
go run cmd/client/main.go
```

## Protobuf
The protobuf definitions for the HelloRequest and HelloResponse messages, as well as the GreetingService service, can be found in the api/hello.proto file. 