package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	// Parse command line arguments
	port := flag.String("port", "8080", "port number")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:"+*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client
	c := pb.NewGreeterClient(conn)

	// Define the name to send in the request
	name := "world"
	if len(flag.Args()) > 0 {
		name = flag.Arg(0)
	}

	// Set a context deadline for the RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello RPC
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("error calling SayHello: %v", err)
	}

	// Print the response
	log.Printf("Response from server: %s", res.Message)
}
