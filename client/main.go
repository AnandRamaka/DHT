package main


import (
	pb "dht/client/pb/inventory"
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// var client_global := pb.NewHashTableClient(conn)

func insertRequest(cl pb.HashTableClient, key string, value string ) {
	request := &pb.InsertRequest{
		Key: key, 
		Value: value,
	}
	
	response, err := cl.InsertValue(context.Background(), request )

	if err != nil {
		log.Fatalf("failed to insert: %v", err)
	}

	fmt.Println(response)

}

func getRequest(cl pb.HashTableClient,  key string) {
	request := &pb.UrlRequest{
		Key: key, 
	}

	response, err := cl.GetValue(context.Background(), request )

	fmt.Println( response )

	if err != nil {
		log.Fatalf("failed to request: %v", err)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHashTableClient(conn)

	log.Println("Context Background:", context.Background())	
	
	insertRequest(client, "2234567890", "test" )

	getRequest(client, "2234567890") 
}


