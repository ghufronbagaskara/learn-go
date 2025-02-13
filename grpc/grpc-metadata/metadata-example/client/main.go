package main

import (
	"context"
	"errors"
	"fmt"
	pb "grpc-metadata/metadata-example/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// create server connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// create client
	client := pb.NewHelloServiceClient(conn)

	// call unary rpc
	callUnary(client)

	// call stream rpc
	callStream(client)
}

func callUnary(client pb.HelloServiceClient) {
	var metadataFromServer, trailerFromServer metadata.MD

	metdataFromClient := metadata.New(map[string]string{"x-client-id": "fast-campus"})

	ctx := metadata.NewOutgoingContext(context.Background(), metdataFromClient)

	resp, err := client.Greet(ctx, nil, grpc.Header(&metadataFromServer), grpc.Trailer(&trailerFromServer))
	if err != nil {
		log.Fatalf("error when calling Greet: %v", err)
	}

	fmt.Println("---Response From Greet---")
	fmt.Println(resp.String())
	fmt.Println()

	fmt.Println("---Metadata From Greet---")
	for key, val := range deduplicateMD(metadataFromServer) {
		fmt.Println(key + ":" + val)
	}
	fmt.Println()

	fmt.Println("---Trailer From Greet---")
	for key, val := range deduplicateMD(trailerFromServer) {
		fmt.Println(key + ":" + val)
	}
	fmt.Println()

}

func callStream(client pb.HelloServiceClient){
	stream, err := client.ServerTime(context.Background(), nil)
	if err != nil {
		log.Fatalf("error when calling ServerTime: %v", err)
	}
	
	fmt.Println("---Response From ServerTime---")
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			break
		}

		if err != nil {
			log.Fatalf("error when reading stream: %v", err)
		}

		fmt.Println(resp.String())
	}

	fmt.Println("---Metadata From ServerTime---")
	metadataFromServer, err := stream.Header()
	if err != nil {
		log.Fatalf("error when reading metadata: %v", err)
	}

	for key, val := range deduplicateMD(metadataFromServer) {
		fmt.Println(key + ":" + val)
	}
	fmt.Println()

	fmt.Println("---Trailer From ServerTime---")
	trailerFromServer := stream.Trailer()
	for key, val := range deduplicateMD(trailerFromServer) {
		fmt.Println(key + ":" + val)
	}
	fmt.Println()

}



func deduplicateMD(metadata metadata.MD) map[string]string {
	metdataReceived := map[string]string{}

	for key, val := range metadata {
		var value string

		for _, v := range val {
			value = fmt.Sprintf("%s%s", value, v)
		}

		metdataReceived[key] = value
	}

	return metdataReceived
}
