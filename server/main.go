package main

import (
	pb "dht/server/pb/inventory"
	"context"
	"log"
	"net"
	"fmt"
	"hash/fnv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"syscall"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % 10000
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

	if ok {
		// Do something
		return &pb.ValueResponse{
			Value: val,
		}, nil
	}
	return &pb.ValueResponse{}, status.Error(400,"Key not found")

}
func CallSuccessor (in *pb.UrlRequest) (*pb.UrlResponse) {
	
	conn, err := grpc.Dial("localhost:" + ports[2], grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	succ_server := pb.NewHashTableClient(conn)

	request := &pb.UrlRequest{
		Key: in.Key,
	}
	
	result, err := succ_server.GetURL(context.Background(), request )
	
	return result

}
func (s *server) GetURL(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {
	fmt.Println("getting url")
	keyHash := int(hash(in.Key))
	pred := serverIds[0]
	answer := ""
	currentId := serverIds[1]
	
	if keyHash == currentId {
		answer = ports[1]		
	} else if keyHash < currentId {
		if keyHash < pred {
			answer = ports[1]
		} else {
			answer = CallSuccessor(in).Url
		}
	} else { //keyHash > serverIds[1]
		if keyHash > pred {
			answer = ports[1]
		} else {
			answer = CallSuccessor(in).Url
		}
	}

	return &pb.UrlResponse{
		Url: answer,
	}, nil
}

func first(n int, _ error) int {
    return n
}

func startNode()(){
	args := os.Args
	serverIds = [3]int{first(strconv.Atoi(args[1])), first(strconv.Atoi(args[3])), first(strconv.Atoi(args[5]))}
	ports = [3]string{args[2], args[4], args[6]} 
	fmt.Println(serverIds)
	fmt.Println(ports)

	listener, err := net.Listen("tcp", "localhost:" + ports[1])
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("started successfully")

	successor := ports[0]
	fmt.Printf("Server started at: " + strconv.Itoa(serverIds[1]) + "  has a successor at: " + successor)

	s := grpc.NewServer()
	
	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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
	defer listener.Close()





	cancelChan := make(chan os.Signal, 1)
    // catch SIGETRM or SIGINTERRUPT
    signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
    go func() {

		fmt.Println("started successfully")
		successor := ports[0]
		fmt.Printf("Server started at: " + strconv.Itoa(serverIds[1]) + "  has a successor at: " + successor)
		s := grpc.NewServer()
		reflection.Register(s)
		pb.RegisterHashTableServer(s, &server{})
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

    }()
    sig := <-cancelChan
    fmt.Println("Caught signal interupt!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	_ = sig
	listener.Close()

	
}
