package main

import (
	"context"
	pb "dht/server/pb/inventory"
	"fmt"
	"hash/fnv"
	"math"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var totalRange uint32 = 10000

func hash(s string) int32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int32(h.Sum32() % totalRange)
}

type server struct {
	pb.UnimplementedHashTableServer
}
type fingerData struct {
	start  int32
	nodeId int32
	port   string
}

var HM = make(map[string]string)
var serverIds [2]int32
var ports [2]string
var id int32
var port string
var isInserting bool
var fingerTable [4]fingerData

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

func containsId(pred int32, currentId int32, keyHash int32) bool {
	if keyHash == currentId || currentId == pred {
		return true
	} else if keyHash < currentId {
		return pred > currentId || keyHash > pred
	}
	//keyHash > currentId
	return keyHash > pred && currentId < pred
}

// keyHash will always be infront of the currentNode
func (s *server) GetClosestFinger(ctx context.Context, in *pb.UrlRequest) (*pb.NodeResponse, error) {
	keyHash := stringToInt32(in.Key)
	for i := 0; i < len(fingerTable); i++ {
		if containsId(serverIds[1], keyHash, fingerTable[i].nodeId) {
			return &pb.NodeResponse{Id: fingerTable[i].nodeId, Url: fingerTable[i].port}, nil
		}
	}
	fmt.Println("SHOULD NOT BE HERE")
	return nil, nil
}
func calculateClosestFinger(keyHash int32) (int32, string) {
	for i := 0; i < len(fingerTable); i++ {
		if containsId(serverIds[1], keyHash, fingerTable[i].nodeId) {
			return fingerTable[i].nodeId, fingerTable[i].port
		}
	}
	fmt.Println("SHOULD NOT BE HERE")
	return 0, ""
}
func findClosestFinger(nodeId int32, url string, keyHash int32) (int32, string) {
	if url == ports[1] {
		return calculateClosestFinger(keyHash)
	}
	conn, node := makeConnection(url)
	request := &pb.UrlRequest{Key: strconv.Itoa(int(keyHash))}
	res, _ := node.GetClosestFinger(context.Background(), request)
	defer conn.Close()
	fmt.Println(res.Id, res.Url)
	return res.Id, res.Url
}

func find_pred(keyHash int32) (int32, string) {
	nodeId := serverIds[1]
	port := ports[1]
	for true {
		succ_data := &pb.NodeResponse{Id: fingerTable[0].nodeId, Url: fingerTable[0].port}
		if port != ports[1] {
			succ_data = callFunction(GetSuccessor, &pb.EmptyRequest{}, port).res.(*pb.NodeResponse)
		}
		if containsId(nodeId, succ_data.Id, keyHash) {
			break
		}
		nodeId, port = findClosestFinger(nodeId, port, keyHash)
	}
	return nodeId, port
}

func (s *server) GetURL(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {

	answerUrl := ""
	var answerId int32 = 0
	if containsId(serverIds[0], serverIds[1], stringToInt32(in.Key)) {
		answerUrl, answerId = ports[1], serverIds[1]
	} else {
		_, predPort := find_pred(stringToInt32(in.Key))
		succ := callFunction(GetSuccessor, &pb.EmptyRequest{}, predPort).res.(*pb.NodeResponse)
		answerId, answerUrl = succ.Id, succ.Url
	}
	if isInserting {
		request := &pb.NeighborUpdate{
			Ports:       ports[1],
			Id:          serverIds[1],
			IsSuccessor: false,
		}

		callFunction(RedistributeKeys, request, fingerTable[0].port)
		isInserting = false
	}
	return &pb.UrlResponse{
		Url: answerUrl,
		Id:  answerId,
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
	fmt.Println("dialing: ", connectionUrl)
	conn, err := grpc.Dial("localhost:"+connectionUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("Make connection error: ", err)
	if err != nil {
		fmt.Println("failed to connect:", err)
	}
	succ_server := pb.NewHashTableClient(conn)
	return conn, succ_server
}

func defaultFingerTable(nodeId int32, nodeUrl string) {
	for i := 0; i < len(fingerTable); i++ {
		fingerTable[i].nodeId = nodeId
		fingerTable[i].port = nodeUrl
	}

	ports[0] = nodeUrl
	ports[1] = nodeUrl
	serverIds[0] = nodeId
	serverIds[1] = nodeId
}

type Functions int64

const (
	GetURL Functions = iota
	GetPredecessor
	ChangeNeighbor
	GetSuccessor
	UpdateRest
	RedistributeKeys
)

type RequestType interface {
	pb.UrlRequest | pb.EmptyRequest
}

// can you fix this error on line 206?
type Response struct {
	res interface{}
}

func voidPrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func callFunction(funcId Functions, request interface{}, port string) *Response {
	conn, otherServer := makeConnection(port)
	defer conn.Close()

	switch funcId {
	case GetURL:
		test := request.(*pb.UrlRequest)
		res, err := otherServer.GetURL(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	case GetPredecessor:
		test := request.(*pb.EmptyRequest)
		res, err := otherServer.GetPredecessor(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	case GetSuccessor:
		test := request.(*pb.EmptyRequest)
		res, err := otherServer.GetSuccessor(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	case ChangeNeighbor:
		test := request.(*pb.NeighborUpdate)
		res, err := otherServer.ChangeNeighbor(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	case UpdateRest:
		test := request.(*pb.FingerUpdate)
		res, err := otherServer.UpdateRest(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	case RedistributeKeys:
		test := request.(*pb.NeighborUpdate)
		res, err := otherServer.RedistributeKeys(context.Background(), test)
		voidPrintError(err)
		return &Response{res: res}
	default:
		fmt.Fprintln(os.Stderr, "DEFAULT SWITCH")
	}

	return &Response{nil} // Add appropriate return statement
}
func initFingerStarts(nodeId int32) {
	for i := 0; i < len(fingerTable); i++ {
		fingerTable[i] = fingerData{int32(uint32(int32(math.Pow(2, float64(i)))+nodeId) % totalRange), 0, ""}
	}
}
func findSuccessor(val int32, port string) (int32, string) {
	request := &pb.UrlRequest{
		Key: strconv.Itoa(int(val)),
	}
	res := callFunction(GetURL, request, port).res

	return res.(*pb.UrlResponse).Id, res.(*pb.UrlResponse).Url
}

func stringToInt32(str string) int32 {
	return int32(first(strconv.Atoi(str)))
}
func (s *server) UpdateRest(ctx context.Context, in *pb.FingerUpdate) (*pb.EmptyResponse, error) {
	index := in.Index
	if containsId(serverIds[1], fingerTable[index].nodeId, in.Val) {
		fingerTable[index].nodeId = in.Val
		fingerTable[index].port = in.Url
		if ports[0] != fingerTable[index].port {
			callFunction(UpdateRest, in, ports[0])
		}
	}
	return &pb.EmptyResponse{}, nil
}
func updateOthers() {
	for i := len(fingerTable) - 1; i >= 0; i-- {
		stepBack := math.Pow(2, float64(i))
		_, port := find_pred((serverIds[1] - int32(stepBack)) % int32(totalRange))
		request := &pb.FingerUpdate{Val: serverIds[1], Url: ports[1], Index: int32(i)}
		callFunction(UpdateRest, request, port)
	}
}

func initFingerTable(nodeId int32, sponsorNodeURL string) {
	fingerTable[0].nodeId, fingerTable[0].port = findSuccessor(nodeId, sponsorNodeURL)

	erequest := &pb.EmptyRequest{}
	res := callFunction(GetPredecessor, erequest, fingerTable[0].port).res
	serverIds[0] = res.(*pb.NodeResponse).Id
	ports[0] = res.(*pb.NodeResponse).Url
	nrequest := &pb.NeighborUpdate{
		Ports:       ports[1],
		Id:          serverIds[1],
		IsSuccessor: false,
	}
	callFunction(ChangeNeighbor, nrequest, fingerTable[0].port)
	//result, err := sponserServer.GetURL(context.Background(), request)
	for i := 1; i < len(fingerTable); i++ {
		if containsId(nodeId, fingerTable[i-1].nodeId, fingerTable[i].start) {
			fingerTable[i].nodeId, fingerTable[i].port = fingerTable[i-1].nodeId, fingerTable[i-1].port
		} else {
			fingerTable[i].nodeId, fingerTable[i].port = findSuccessor(fingerTable[i].start, sponsorNodeURL)
		}
	}
	PrintFingers()
	updateOthers()
}

func insertNode(nodeId int32, nodeUrl string, sponsorNodeId int32, sponsorNodeURL string) {
	initFingerStarts(nodeId)
	if sponsorNodeId < 0 {
		isInserting = false
		defaultFingerTable(nodeId, nodeUrl)
	} else {
		isInserting = true
		initFingerTable(nodeId, sponsorNodeURL)
	}
	startNode()

}

func (s *server) ChangeNeighbor(ctx context.Context, in *pb.NeighborUpdate) (*pb.NodeResponse, error) {
	if in.IsSuccessor {
		fingerTable[0].nodeId = in.Id
		fingerTable[0].port = in.Ports
	} else {
		serverIds[0] = int32(in.Id)
		ports[0] = in.Ports
	}

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
func PrintFingers() {
	fmt.Println("FINGER TABLE")
	for i := 0; i < len(fingerTable); i++ {
		fmt.Println(fingerTable[i].start, " ", fingerTable[i].nodeId, "  ", fingerTable[i].port)
	}
}
func (s *server) GetSuccessor(ctx context.Context, in *pb.EmptyRequest) (*pb.NodeResponse, error) {
	//PrintFingers()
	return &pb.NodeResponse{
		Url: fingerTable[0].port,
		Id:  fingerTable[0].nodeId,
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
	for key, element := range HM {
		if hash(key) <= in.Id {
			newServer.InsertValue(context.Background(), &pb.InsertRequest{Key: key, Value: element})
		}
	}
	for key, _ := range HM {
		if hash(key) <= in.Id {
			delete(HM, key)
		}
	}
	conn.Close()
	return &pb.EmptyResponse{}, nil
}

func main() {
	args := os.Args
	fmt.Println("ARGS  ", args)

	// args = [nodeId, nodeUrl, sponsorNodeId, sponsorNodeURL ]
	fmt.Println(args)
	serverIds[1] = stringToInt32(args[1])
	ports[1] = args[2]
	insertNode(stringToInt32(args[1]), args[2], stringToInt32(args[3]), args[4])

	fmt.Println(serverIds)
	fmt.Println(ports)
}
