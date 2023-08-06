package main

import (
	"log"

	pb "github.com/saifuljnu/gRPC-Unary/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

func main() {
	// Step 1: Dial the gRPC server (insecure connection)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// Step 2: Create the gRPC service client
	c := pb.NewGreetServiceClient(conn)

	// Step 3: Call the gRPC service method
	doGreet(c)
}
