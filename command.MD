# install protocol buffer compiler
1) sudo apt-get install protobuf-compiler
2) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
3) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# manually generate code
protoc -Igreet/proto --go_opt=module=github.com/renatospaka/grpc-calculator --go_out=. --go-grpc_opt=module=github.com/renatospaka/grpc-calculator --go-grpc_out=. greet/proto/*.proto
protoc -I<PATH PROTO FILE> --go_out=<PATH TO COMPILED FILES> --go-grpc_out=<PATH TO COMPILED FILES> <NAME (AND PATH) OF THE PROTO FILE>

# generate and remove code with Makefile
make greet // creates pb files for greet folder
make clean_greet // delete pb files from greet folder
make help