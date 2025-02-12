package main

import (
	"context"
	pb "grpc-example/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "Azzachra",
	})

	if err != nil {
		log.Printf("an error occured: %v", err)
		return
	}

	log.Printf("Message from server: %s", r.GetMessage())
}
