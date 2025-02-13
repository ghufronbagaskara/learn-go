package main

import (
	"context"
	"fmt"
	pb "grpc-metadata/metadata-example/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Greet(ctx context.Context, _ *emptypb.Empty) (*pb.GreetResponse, error) {
	currentTime := time.Now()

	defer func() {
		serverTrailer := metadata.New(map[string]string{"x-process-time": fmt.Sprintf("%dms", time.Since(currentTime).Microseconds())})
		grpc.SetTrailer(ctx, serverTrailer)
	}()

	response := make(map[string]string)

	meta, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, val := range meta {
			var value string

			for _, v := range val {
				value += fmt.Sprintf("%s%s", value, v)
			}

			response[key] = value
		}
	}

	// create new metadata
	serverMetadata := metadata.New(map[string]string{"x-server-rpc": "greet-rpc"})
	grpc.SendHeader(ctx, serverMetadata)

	// kirim response
	return &pb.GreetResponse{
		Metadata: response,
	}, nil
}

func (s *server) ServerTime(_ *emptypb.Empty, stream pb.HelloService_ServerTimeServer) error {
	return nil
}
