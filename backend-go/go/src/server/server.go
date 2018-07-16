package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

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
		return nil, status.Error(codes.Internal, "Error response")
	}

	return &pb.EchoResponse{
		Message: in.Message,
	}, nil
}

func (s *echoService) ServerStreamingEcho(in *pb.ServerStreamingEchoRequest, stream pb.EchoService_ServerStreamingEchoServer) error {
	// Ensure conformity with PROTOCOL-WEB, see
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2
	stream.SendHeader(metadata.Pairs("accept", "application/grpc-web-text"))

	if in.Message == "error" {
		return status.Error(codes.Internal, "Error response #2")
	}

	for index := 0; index < int(in.GetMessageCount()); index++ {
		time.Sleep(time.Duration(in.GetMessageInterval()) * time.Millisecond)
		msg := &pb.ServerStreamingEchoResponse{
			Message: in.GetMessage(),
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
