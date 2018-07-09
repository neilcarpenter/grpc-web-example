package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "proto/echo"
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
	pb.RegisterEchoServiceServer(gs, &echoService{})

	log.Printf("starting grpc on :%d\n", *grpcPort)

	gs.Serve(lis)
}

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	// Ensure conformity with PROTOCOL-WEB, see
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2
	grpc.SendHeader(ctx, metadata.Pairs("accept", "application/grpc-web-text"))

	if in.Message == "error" {
		return nil, status.Error(codes.Internal, "pb error response")
	}

	return &pb.EchoResponse{
		Message: in.Message,
	}, nil
}
