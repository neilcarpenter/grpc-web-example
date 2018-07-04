package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Failed to retrieve metadata from incoming request")
		return nil, status.Error(codes.Internal, "Failed to retrieve metadata from incoming request")
	}

	// Copy client metadata to response, this is to ensure gRPC-web filter in Envoy is applied correctly
	grpc.SendHeader(ctx, md)

	if in.Message == "error" {
		return nil, status.Error(codes.Internal, "pb error response")
	}

	return &pb.EchoResponse{
		Message: in.Message,
	}, nil
}
