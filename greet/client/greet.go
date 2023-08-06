package main

import (
	"context"
	"log"

	pb "github.com/saifuljnu/gRPC-Unary/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	// Make a gRPC unary RPC call to the server's Greet method
	r, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Saiful",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
