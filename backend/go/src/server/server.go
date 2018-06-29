package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	example "proto/example"
)

var (
	grpcPort = flag.Int("port", 9091, "grpc port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	example.RegisterEchoServiceServer(gs, &echoService{})

	log.Printf("starting grpc on :%d\n", *grpcPort)

	gs.Serve(lis)
}

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, in *example.EchoRequest) (*example.EchoResponse, error) {
	return &example.EchoResponse{
		Message: in.Message,
	}, nil
}
