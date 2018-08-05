package main

import (
    "google.golang.org/grpc"

    pb "github.com/cclin81922/k8s-microsvc-quickstart/pkg/pb/helloworld"
    server "github.com/cclin81922/k8s-microsvc-quickstart/pkg/server/helloworld"
)

func registerHelloworld(s *grpc.Server, client interface{}) {
	pb.RegisterGreeterServer(s, &server.Server{client})
}
