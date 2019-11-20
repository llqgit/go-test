package main

import (
	"context"
	"fmt"
	rpc "github.com/llqgit/go-test/grpc/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":53333"

type RpcService struct {
}

func (s *RpcService) SayHello(ctx context.Context, request *rpc.HelloRequest) (*rpc.HelloReply, error) {
	return &rpc.HelloReply{Message: "Hello - " + request.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	rpc.RegisterHelloServer(s, &RpcService{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("start server")
}
