package main

import (
	"context"
	"fmt"
	pb "grpc-example/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Received :", req.GetName())
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s", req.GetName())}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server is listening at %v :", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
