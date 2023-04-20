package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/renatospaka/grpc-calculator/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Renato"},
		{FirstName: "Debora"},
		{FirstName: "Mara"},
		{FirstName: "Jana"},
		{FirstName: "Pedrão"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
	
			time.Sleep(500 * time.Millisecond)
		}
	
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
	
			if err != nil {
				log.Fatalf("Error while reading client stream: %v\n", err)
				break
			}

			log.Printf("GreetEveryone: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}