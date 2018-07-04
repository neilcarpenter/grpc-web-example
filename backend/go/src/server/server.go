package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Failed to retrieve metadata from incoming request")
		return nil, status.Error(codes.Internal, "Failed to retrieve metadata from incoming request")
	}

	// Copy client metadata to response, this is to ensure gRPC-web filter in Envoy is applied correctly
	grpc.SendHeader(ctx, md)

	if in.Message == "error" {
		return nil, status.Error(codes.Internal, "Example error response")
	}

	return &example.EchoResponse{
		Message: in.Message,
	}, nil
}

func (s *echoService) ServerStreamingEcho(in *example.ServerStreamingEchoRequest, stream example.EchoService_ServerStreamingEchoServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Printf("Failed to retrieve metadata from incoming request")
		return status.Error(codes.Internal, "Failed to retrieve metadata from incoming request")
	}

	stream.SendHeader(md)

	for index := 0; index < int(in.GetMessageCount()); index++ {
		time.Sleep(time.Duration(in.GetMessageInterval()) * time.Millisecond)
		msg := &example.ServerStreamingEchoResponse{
			Message: in.GetMessage(),
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
