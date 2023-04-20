package main

import (
	"io"
	"log"

	pb "github.com/renatospaka/grpc-calculator/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function invoked")

	var sum int32 = 0
	count := 00

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the client stream: %v\n", err)
		}
		
		log.Printf("Receiving number: %d\n", req.Number)
		sum += req.Number
		count++
	}
}
