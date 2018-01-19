#GO protoc instructions
running from 'assets' folder set GOPATH and PATH: `GOPATH=$(cd ../../../ && pwd) && PATH=$GOPATH/bin:$PATH`</br>
install 'protoc-gen-go' if needed: `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`</br>
generate GO source files: `protoc --go_out=../pkg/protobuf *.proto`</br>

install grpc plugin: `go get -u google.golang.org/grpc`
protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide