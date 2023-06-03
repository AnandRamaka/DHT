package main


import (
	pb "dht/client/pb/inventory"
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"strings"
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
	ports_file := "../ports.txt"
	content, err := ioutil.ReadFile(ports_file)
	
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var portList []string
	for _, line := range lines {
		portList = append(portList, line)
	}

	conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHashTableClient(conn)

	log.Println("Context Background:", context.Background())	
	
	insertRequest(client, "2234567890", "test" )
	getRequest(client, "2234567890") 
}


