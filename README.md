Setup
Linux
GOROOT=/usr/lib/go
GOPATH=/usr/aryanj/go
MAC export GOROOT=/usr/local/go
export GOPATH=/Users/aryanjoshi/go
Common
all code should go under GOPATH/src
clone repo
go mod init bookshop/server
protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
from server folder run
go install bookshop/server
the server is ready to start
