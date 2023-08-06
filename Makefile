# Makefile

# Generate gRPC code from greet.proto
run:
	protoc -I greet/proto --go_out=. --go_opt=module=github.com/saifuljnu/gRPC-Unary --go-grpc_out=. --go-grpc_opt=module=github.com/saifuljnu/gRPC-Unary greet/proto/greet.proto

# Default rule: Build both server and client
all: server client

# Build the server
server:
	go build -o bin/greet/server ./greet/server

# Build the client
client:
	go build -o bin/greet/client ./greet/client
