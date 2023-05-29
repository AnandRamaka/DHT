package main

import (
	pb "dht/server/pb/inventory"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHashTableServer
}

// func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
// 	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
// 	return &pb.GetBookListResponse{
// 		Books: getSampleBooks(),
// 	}, nil
// }
 
func (s *server) GetValue(ctx context.Context, in *pb.UrlRequest) (*pb.ValueResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.ValueResponse{
				Value: "This is a value",
			}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func getSampleBooks() []*pb.Book {
// 	sampleBooks := []*pb.Book{
// 		{
// 			Title:     "The Hitchhiker's Guide to the Galaxy",
// 			Author:    "Douglas Adams",
// 			PageCount: 42,
// 		},
// 		{
// 			Title:     "The Lord of the Rings",
// 			Author:    "J.R.R. Tolkien",
// 			PageCount: 1234,
// 		},
// 	}
// 	return sampleBooks
// }
