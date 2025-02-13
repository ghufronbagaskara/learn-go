package main

import (
	"context"
	"fmt"
	pb "grpc-metadata/metadata-example/proto"
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
