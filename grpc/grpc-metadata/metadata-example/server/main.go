package main

import (
	"log"
	"net"

	pb "grpc-metadata/metadata-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterHelloServiceServer(srv, &server{})

	reflection.Register(srv)
	log.Printf("server listening at: %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
