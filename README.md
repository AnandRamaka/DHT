Implementation of the CHORD Protocal, which defines a peer-to-peer key value store.

**Setup**

**Linux**

_GOROOT=/usr/lib/go_

_GOPATH=/usr/<username>/go_



**Mac**

_export GOROOT=/usr/local/go_

_export GOPATH=/Users/<username>/go_



**Common**

all GO code should go under GOPATH/src

clone repo

_go mod init bookshop/server

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=._
^ This protoc command generates programs that implement functionality defined in the proto files


from server folder run:

_go install bookshop/server_

the server is ready to start

The python script launches multiple processes of the server and cleans all processes

