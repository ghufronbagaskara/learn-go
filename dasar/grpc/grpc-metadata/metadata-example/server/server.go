package main

import (
	"context"
	"fmt"
	pb "grpc-metadata/metadata-example/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

// Unary 
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


// Stream
func (s *server) ServerTime(_ *emptypb.Empty, stream pb.HelloService_ServerTimeServer) error {
	currentTime := time.Now()

	defer func(){
		serverTrailer := metadata.New(map[string]string{"x-process-time": fmt.Sprintf("%dms", time.Since(currentTime).Milliseconds())})
		stream.SetTrailer(serverTrailer)
	}()
	
	serverMetadata := metadata.New(map[string]string{"x-server-rpc": "server-time-rpc"})
	stream.SendHeader(serverMetadata)

	for time.Since(currentTime) < 3 *time.Second {
		select {
			case <- stream.Context().Done() : 
				return stream.Context().Err()
			case <- time.After(1 * time.Second) :
				if err := stream.Send(&pb.ServerTimeResponse{
					CurrentTime: timestamppb.Now(),
				}); err != nil {
					return err
				}
		}
	}

	return nil
}


