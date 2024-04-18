package main

import (
    "context"
    "flag"
    "log"
    "google.golang.org/grpc"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
    port := flag.String("port", "8080", "port number")
    flag.Parse()

    conn, err := grpc.Dial("localhost:"+*port, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("could not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    name := "world"
    if len(flag.Args()) > 0 {
        name = flag.Arg(0)
    }
    res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("error calling SayHello: %v", err)
    }
    log.Printf("Response from server: %s", res.Message)
}
