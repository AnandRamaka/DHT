package main

import (
	pb "dht/server/pb/inventory"
	"context"
	"log"
	"net"
	"os"
	"fmt"
	"hash/fnv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
	"strconv"
	"google.golang.org/grpc/status"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type server struct {
	pb.UnimplementedHashTableServer
}

var HM = make(map[string]string)
var serverIds [3]int
var ports [3]string
func (s *server) InsertValue(ctx context.Context, in *pb.InsertRequest) (*pb.Status, error) {
	HM[in.Key] = in.Value


	fmt.Println(serverIds[0])

	return &pb.Status{
		Result: "Success",
	}, nil
}

func (s *server) GetValue(ctx context.Context, in *pb.UrlRequest) (*pb.ValueResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	val, ok := HM[in.Key]
	// If the key exists
	fmt.Println("Got Request for key" + in.Key)
	if ok {
		// Do something
		return &pb.ValueResponse{
			Value: val,
		}, nil
	}
	return &pb.ValueResponse{}, status.Error(400,"Key not found")
}
// func (s *server) GetURL(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {
	
// }
func first(n int, _ error) int {
    return n
}

func main() {
	args := os.Args

	serverIds = [3]int{first(strconv.Atoi(args[1])), first(strconv.Atoi(args[3])), first(strconv.Atoi(args[5]))}
	ports = [3]string{args[2], args[4], args[6]} 
	fmt.Println(serverIds)
	fmt.Println(ports)
	listener, err := net.Listen("tcp", "localhost:" + ports[1])
	if err != nil {
		panic(err)
	}
	
	//Waiting for other servers to start up 
	time.Sleep(5 * time.Second)

	successor := args[2]
	fmt.Printf("Server started at: " + args[1] + "  has a successor at: " + successor)

	s := grpc.NewServer()
	
	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
