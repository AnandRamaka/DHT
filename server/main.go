package main

import (
	pb "dht/server/pb/inventory"
	"context"
	"log"
	"net"
	"os"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//args[1]

	s := grpc.NewServer()
	
	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
