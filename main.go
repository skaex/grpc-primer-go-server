package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/skaex/grpc-primer-go-server/messages"
	"google.golang.org/grpc"
)

const port = ":50000"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	messages.RegisterHelloServiceServer(s, &server{})

	s.Serve(listen)
}
