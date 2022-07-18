# install protocol buffer compiler
1) sudo apt-get install protobuf-compiler
2) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
3) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# manually generate code
protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto
protoc -I<PATH PROTO FILE> --go_out=<PATH TO COMPILED FILES> --go-grpc_out=<PATH TO COMPILED FILES> <NAME (AND PATH) OF THE PROTO FILE>
