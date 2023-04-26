package main

import (
	"log"

	pb "github.com/renatospaka/grpc-calculator/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	opts := []grpc.DialOption{}
	tls := true
	if tls {
		certfile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certfile, "")
		if err != nil {
			log.Fatalf("Filed to load CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5 * time.Second)
	// doGreetWithDeadline(c, 1 * time.Second)
}
