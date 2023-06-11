package main

import (
	"context"
	pb "dht/server/pb/inventory"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func hash(s string) int32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int32(h.Sum32() % 10000)
}

type server struct {
	pb.UnimplementedHashTableServer
}

var HM = make(map[string]string)
var serverIds [3]int32
var ports [3]string
var id int32
var port string

func (s *server) InsertValue(ctx context.Context, in *pb.InsertRequest) (*pb.Status, error) {
	HM[in.Key] = in.Value

	fmt.Println(in.Key, in.Value)

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
	return &pb.ValueResponse{}, status.Error(400, "Key not found")
}

func CallSuccessor(in *pb.UrlRequest) *pb.UrlResponse {

	conn, err := grpc.Dial("localhost:"+ports[2], grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	succ_server := pb.NewHashTableClient(conn)

	request := &pb.UrlRequest{
		Key: in.Key,
	}

	result, err := succ_server.GetURL(context.Background(), request)

	return result

}
func (s *server) GetURL(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {
	fmt.Println("getting url")
	keyHash := int32(hash(in.Key))
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

func startNode() {
	args := os.Args

	serverIds = [3]int32{int32(first(strconv.Atoi(args[1]))), int32(first(strconv.Atoi(args[3]))), int32(first(strconv.Atoi(args[5])))}
	ports = [3]string{args[2], args[4], args[6]}
	fmt.Println(serverIds)
	fmt.Println(ports)
	id = serverIds[1]
	listener, err := net.Listen("tcp", "localhost:"+ports[1])
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("started successfully")

	successor := ports[0]
	fmt.Println("Server started at: " + strconv.Itoa(int(serverIds[1])) + "  has a successor at: " + successor)

	s := grpc.NewServer()

	reflection.Register(s)
	pb.RegisterHashTableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func makeConnection(connectionUrl string) (*grpc.ClientConn, pb.HashTableClient) {
	conn, err := grpc.Dial("localhost:"+connectionUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	succ_server := pb.NewHashTableClient(conn)
	return conn, succ_server
}

func insertNode(nodeId int, nodeUrl string, sponsorNodeURL string) {
	// 1, Get the node before from the sponsor node
	// Modify successor of node before
	// Modify predecessor of the node after the node before
	// Redistribute the keys from the after node to the newly inserted node

	conn, sponserServer := makeConnection(nodeUrl)
	conn.Close()

	request := &pb.UrlRequest{
		Key: string(nodeId),
	}

	result, _ := sponserServer.GetURL(context.Background(), request)
	newSuccessor := result.Url

	conn2, succ_server := makeConnection(newSuccessor)

	predUrl, _ := succ_server.GetPredecessor(context.Background(), &pb.EmptyRequest{})
	conn2.Close()
	conn3, predServer := makeConnection(predUrl.Url)

	request2 := &pb.NeighborUpdate{
		Ports:       result.Url,
		Id:          int32(nodeId),
		IsSuccessor: true,
	}

	predServer.ChangeNeighbor(context.Background(), request2)
	conn3.Close()

	//GetPredecessor
}

func (s *server) ChangeNeighbor(ctx context.Context, in *pb.NeighborUpdate) (*pb.EmptyResponse, error) {
	int_val := 0
	if in.IsSuccessor {
		int_val = 1
	}

	serverIds[2*int_val] = int32(in.Id)
	ports[2*int_val] = in.Ports

	return &pb.EmptyResponse{}, nil
}

func (s *server) GetPredecessor(ctx context.Context, in *pb.EmptyRequest) (*pb.NodeResponse, error) {
	return &pb.NodeResponse{
		Url: ports[0],
		Id:  serverIds[0],
	}, nil
}

func (s *server) RedistributeKeys(ctx context.Context, in *pb.NeighborUpdate) (*pb.EmptyResponse, error) {
	for key, element := range HM {
		if hash(key) <= in.Id {
			conn, newServer := makeConnection(in.Ports)
			newServer.InsertValue(context.Background(), &pb.InsertRequest{Key: key, Value: element})
			conn.Close()
		}
	}
	return &pb.EmptyResponse{}, nil
}

func main() {
	args := os.Args
	if len(args) == 7 {
		startNode()
	} else {
		// args = [nodeId, nodeUrl, sponsorNodeURL ]
		insertNode(first(strconv.Atoi(args[1])), args[2], args[3])
	}
	fmt.Println(serverIds)
	fmt.Println(ports)

	startNode()
}
