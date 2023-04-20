package main

import (
	"context"
	"log"
	"time"

	pb "github.com/renatospaka/grpc-calculator/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	
	reqs := []*pb.GreetRequest{
		{FirstName: "Renato"},
		{FirstName: "Debora"},
		{FirstName: "Mara"},
		{FirstName: "Jana"},
		{FirstName: "Pedrão"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)

		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reading the stream: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
