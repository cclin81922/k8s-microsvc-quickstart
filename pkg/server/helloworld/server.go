package helloworld

import (
    "log"

    "golang.org/x/net/context"

    pb "github.com/cclin81922/k8s-microsvc-quickstart/pkg/pb/helloworld"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{
    Client interface{}
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

    log.Printf("Client object %v", s.Client)

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}


