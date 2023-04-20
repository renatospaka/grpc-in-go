package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/renatospaka/grpc-calculator/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 31}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})

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
				log.Fatalf("Error while reading server stream: %v\n", err)
				break
			}

			log.Printf("New maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
