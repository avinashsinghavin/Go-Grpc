package main

import (
	"context"
	"log"
	"proto-go/calculator/proto"
)

func doCalculate(c proto.CalculatorServiceClient) {
	log.Printf("calculating value as per request")
	res, err := c.Calculate(context.Background(), &proto.CalculationRequest{
		X:         20,
		Y:         12,
		Operation: "+",
	})

	if err != nil {
		log.Fatalf("could not make calculation %v\n", err)
	}

	log.Printf("Calculating Result is %v \n", res.Result)
}
