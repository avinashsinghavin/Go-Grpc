package main

import (
	"context"
	"io"
	"log"
	"proto-go/calculator/proto"
)

func findPrimary(c proto.CalculatorServiceClient) {
	log.Print("started client to find prime composition")
	req := &proto.NumberRequest{
		Number: 120,
	}

	stream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Printf("Error while calculating prime decompositions %v\n", err)
	}

	for true {
		mes, err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Printf("Error while reading the stream of prime no %v\n", err)
		}

		if err == io.EOF {
			log.Printf("Break as all the required field got %v\n", err)
			break
		}
		log.Printf(" %v ", mes)
	}
	log.Printf("\nSuccessfully got all the Prime decomposition no of %v\n", req.Number)
}
