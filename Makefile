gen:
	protoc --proto_path=proto proto/*.proto --go_out=.
clean:
	rm -f pb/*.pb.go
run:
	go run main.go
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1