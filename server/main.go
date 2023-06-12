package main

import (
	"context"
	pb "dht/server/pb/inventory"
	"fmt"
	"hash/fnv"
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
var isInserting bool

func (s *server) InsertValue(ctx context.Context, in *pb.InsertRequest) (*pb.Status, error) {
	HM[in.Key] = in.Value

	fmt.Println("Just inserted ", in.Key, in.Value)

	return &pb.Status{
		Result: "Success",
	}, nil
}

func (s *server) GetValue(ctx context.Context, in *pb.UrlRequest) (*pb.ValueResponse, error) {
	fmt.Println("Received request: ", in.ProtoReflect().Descriptor().FullName())
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
		fmt.Println("failed to connect: ", err)
	}
	defer conn.Close()
	succ_server := pb.NewHashTableClient(conn)

	request := &pb.UrlRequest{
		Key: in.Key,
	}

	result, err := succ_server.GetURL(context.Background(), request)
	fmt.Println("The key is in", result.Url)
	return result

}
func (s *server) GetURL(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {
	fmt.Println("getting url for key: ", in.Key)
	keyHash := int32(hash(in.Key))
	pred := serverIds[0]
	answer := ""
	currentId := serverIds[1]
	fmt.Println(pred, currentId, keyHash)
	if keyHash == currentId || currentId == pred {
		fmt.Println("case 1")
		answer = ports[1]
	} else if keyHash < currentId {
		if pred > currentId || keyHash > pred {
			fmt.Println("case 2")
			answer = ports[1]
		} else {
			answer = CallSuccessor(in).Url
		}
	} else { //keyHash > currentId
		if keyHash > pred && currentId < pred {
			fmt.Println("case 3")
			answer = ports[1]
		} else {
			answer = CallSuccessor(in).Url
		}
	}

	if isInserting {
		conn, succ := makeConnection(ports[2])

		request := &pb.NeighborUpdate{
			Ports:       ports[1],
			Id:          serverIds[1],
			IsSuccessor: true,
		}

		succ.RedistributeKeys(context.Background(), request)
		isInserting = false
		conn.Close()
	}

	return &pb.UrlResponse{
		Url: answer,
	}, nil
}

func first(n int, _ error) int {
	return n
}

func startNode() {

	id = serverIds[1]
	listener, err := net.Listen("tcp", "localhost:"+ports[1])
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println(serverIds)
	fmt.Println(ports)
	fmt.Println("started successfully")

	successor := ports[0]

	fmt.Println("Server Id: ", serverIds[1])
	fmt.Println("Server started at: " + strconv.Itoa(int(serverIds[1])) + "  has a successor at: " + successor)

	s := grpc.NewServer()

	reflection.Register(s)

	pb.RegisterHashTableServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		fmt.Println("failed to serve: ", err)
	}

	fmt.Println("server running")

}

func makeConnection(connectionUrl string) (*grpc.ClientConn, pb.HashTableClient) {
	conn, err := grpc.Dial("localhost:"+connectionUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("Make connection error: ", err)
	if err != nil {
		fmt.Println("failed to connect:", err)
	}
	succ_server := pb.NewHashTableClient(conn)
	return conn, succ_server
}

func insertNode(nodeId int, nodeUrl string, sponsorNodeURL string) {
	// 1, Get the node before from the sponsor node
	// Modify successor of node before
	// Modify predecessor of the node after the node before
	// Redistribute the keys from the after node to the newly inserted node

	conn, sponserServer := makeConnection(sponsorNodeURL)

	request := &pb.UrlRequest{
		Key: strconv.Itoa(nodeId),
	}

	fmt.Println(nodeId, nodeUrl)

	result, err := sponserServer.GetURL(context.Background(), request)
	fmt.Println(err)

	newSuccessor := result.Url

	// fmt.Println(result.Url)
	conn.Close()
	conn2, succ_server := makeConnection(newSuccessor)

	predData, err := succ_server.GetPredecessor(context.Background(), &pb.EmptyRequest{})
	conn2.Close()
	conn3, predServer := makeConnection(predData.Url)

	request2 := &pb.NeighborUpdate{
		Ports:       nodeUrl,
		Id:          int32(nodeId),
		IsSuccessor: false,
	}

	_, err = predServer.ChangeNeighbor(context.Background(), request2)
	conn3.Close()

	conn5, predServer := makeConnection(predData.Url)

	request2 = &pb.NeighborUpdate{
		Ports:       nodeUrl,
		Id:          int32(nodeId),
		IsSuccessor: true,
	}

	predServer.ChangeNeighbor(context.Background(), request2)
	conn5.Close()

	fmt.Println("inserted server succesfully")

	conn4, succ := makeConnection(result.Url)
	data, err := succ.GetNodeData(context.Background(), &pb.EmptyRequest{})
	conn4.Close()

	serverIds[0] = predData.Id
	serverIds[1] = int32(nodeId)
	serverIds[2] = data.Id

	ports[0] = predData.Url
	ports[1] = nodeUrl
	ports[2] = result.Url

	startNode()

}

func (s *server) ChangeNeighbor(ctx context.Context, in *pb.NeighborUpdate) (*pb.NodeResponse, error) {
	int_val := 0
	if in.IsSuccessor {
		int_val = 1
		fmt.Println("updated succ")
	} else {
		fmt.Println("updated pred")
	}

	serverIds[2*int_val] = int32(in.Id)
	ports[2*int_val] = in.Ports
	fmt.Println(serverIds)
	fmt.Println(ports)
	return &pb.NodeResponse{
		Url: ports[0],
		Id:  serverIds[0],
	}, nil
}

func (s *server) GetPredecessor(ctx context.Context, in *pb.EmptyRequest) (*pb.NodeResponse, error) {
	return &pb.NodeResponse{
		Url: ports[0],
		Id:  serverIds[0],
	}, nil
}

func (s *server) GetNodeData(ctx context.Context, in *pb.EmptyRequest) (*pb.NodeResponse, error) {
	return &pb.NodeResponse{
		Url: ports[1],
		Id:  serverIds[1],
	}, nil
}

func (s *server) RedistributeKeys(ctx context.Context, in *pb.NeighborUpdate) (*pb.EmptyResponse, error) {
	conn, newServer := makeConnection(in.Ports)
	fmt.Println("REDISTRIBUTING")
	for key, element := range HM {
		if hash(key) <= in.Id {
			fmt.Println("redistribute", key)
			newServer.InsertValue(context.Background(), &pb.InsertRequest{Key: key, Value: element})
		}
	}
	conn.Close()
	return &pb.EmptyResponse{}, nil
}

func main() {
	args := os.Args
	fmt.Println("ARGS  ", args)
	if len(args) == 7 {
		fmt.Println("initial server")
		isInserting = false
		serverIds = [3]int32{int32(first(strconv.Atoi(args[1]))), int32(first(strconv.Atoi(args[3]))), int32(first(strconv.Atoi(args[5])))}
		ports = [3]string{args[2], args[4], args[6]}
		startNode()
	} else {
		isInserting = true
		// args = [nodeId, nodeUrl, sponsorNodeURL ]
		insertNode(first(strconv.Atoi(args[1])), args[2], args[3])
	}
	fmt.Println(serverIds)
	fmt.Println(ports)
}
