package main

import (
	pb "dht/client/pb/inventory"
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHashTableClient(conn)
	bookList, err := client.GetValue(context.Background(), &pb.UrlRequest{} )
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("Test value: %v", bookList)
}
