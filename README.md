**Setup**

**Linux**

_GOROOT=/usr/lib/go_

_GOPATH=/usr/aryanj/go_



**Mac**

_export GOROOT=/usr/local/go_

_export GOPATH=/Users/aryanjoshi/go_



**Common**

all code should go under GOPATH/src

clone repo

_go mod init bookshop/server

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=._

from server folder run:

_go install bookshop/server_

the server is ready to start
