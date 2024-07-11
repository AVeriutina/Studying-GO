package main

import (
	"awesomeProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Printf("Received: %s\n", req.Name)

	return &proto.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
