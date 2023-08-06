package main

import (
	"log"
	"net"

	pb "github.com/saifuljnu/gRPC-Unary/greet/proto" // Corrected import path for the protocol buffer package
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// Now we have to create gRPC server object
// to implement all the rpc endpoints that we will define in our greet proto
// Step 2:
type Server struct {
	// We need to reference the server that was generated for us by gRPC plugin for protocol buffer compiler
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Step 3:
	// Creating our gRPC instance
	s := grpc.NewServer()

	// Step 4:
	// Register the GreetServiceServer
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
