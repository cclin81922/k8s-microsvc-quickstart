package main

import (
	"log"

	"golang.org/x/net/context"
    "google.golang.org/grpc"

	pb "github.com/cclin81922/k8s-microsvc-quickstart/pkg/pb/helloworld"
)

func testHelloworld(conn *grpc.ClientConn, ctx context.Context) {
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
