package main

import (
	"context"
	pb "dht/client/pb/inventory"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeConnection(connectionUrl string) (*grpc.ClientConn, pb.HashTableClient) {
	conn, err := grpc.Dial("localhost:"+connectionUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("failed to connect:", err)
	}
	succ_server := pb.NewHashTableClient(conn)
	return conn, succ_server
}

func getUrlRequest(cl pb.HashTableClient, key string) string {
	request := &pb.UrlRequest{
		Key: key,
	}
	response, err := cl.GetURL(context.Background(), request)
	if err != nil {
		fmt.Println("getUrl failed")
	}
	return response.Url
}
func insertRequest(cl pb.HashTableClient, key string, value string) {
	destinationUrl := getUrlRequest(cl, key)
	request := &pb.InsertRequest{
		Key:   key,
		Value: value,
	}
	conn, destServer := makeConnection(destinationUrl)

	_, err := destServer.InsertValue(context.Background(), request)

	if err != nil {
		fmt.Println("failed to insert: ", err)
	}

	conn.Close()
}

func getRequest(cl pb.HashTableClient, key string) {
	destinationUrl := getUrlRequest(cl, key)
	request := &pb.UrlRequest{
		Key: key,
	}
	conn, destServer := makeConnection(destinationUrl)

	response, err := destServer.GetValue(context.Background(), request)

	if err != nil {
		log.Fatalf("failed to request: %v", err)
		fmt.Println("getRequest failed")
	} else {
		fmt.Println("Response ", key, ": ", response.Value)

	}
	conn.Close()
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

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHashTableClient(conn)

	// insertRequest(client, "3054", "10")
	// insertRequest(client, "4000", "12")
	// insertRequest(client, "0", "stupid")
	// insertRequest(client, "11", "lala")
	// insertRequest(client, "49", "juice")
	// insertRequest(client, "random", "112")

	// getRequest(client, "3054")
	// getRequest(client, "4000")
	// getRequest(client, "0")
	// getRequest(client, "11")
	// getRequest(client, "49")
	// getRequest(client, "random")
	insertRequest(client, "9990", "10")
}
