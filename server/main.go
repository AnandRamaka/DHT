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

func (s *server) InsertValue(ctx context.Context, in *pb.InsertRequest) (*pb.Status, error) {
	HM[in.Key] = in.Value

	return &pb.Status{
		Result: "Success",
	}, nil
}

func (s *server) GetValue(ctx context.Context, in *pb.UrlRequest) (*pb.ValueResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())

	return &pb.ValueResponse{
				Value: HM[in.Key],
			}, nil
}

func main() {
	args := os.Args
	listener, err := net.Listen("tcp", args[1])
	if err != nil {
		panic(err)
	}

	successor := args[2]
	fmt.Printf("Server started at: " + args[1] + "  has a successor at: " + successor)

	s := grpc.NewServer()
	
	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
