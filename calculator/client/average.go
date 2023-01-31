package main

import (
	"context"
	"fmt"
	"proto-go/calculator/proto"
	"time"
)

func doAverage(client proto.CalculatorServiceClient) {
	fmt.Println("Client site streaming")
	numberList := []*proto.AverageRequest{
		{
			Number: 10,
		}, {
			Number: 20,
		}, {
			Number: 30,
		}, {
			Number: 40,
		}, {
			Number: 50,
		}, {
			Number: 60,
		}, {
			Number: 70,
		}, {
			Number: 80,
		}, {
			Number: 90,
		},
	}

	stream, err := client.AverageClientStreaming(context.Background())
	if err != nil {
		fmt.Printf("Error while calling Average Server Streaming %v \n", err)
	}

	for _, req := range numberList {
		fmt.Printf("send no: %d\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Error while receiving response from client Streaming %v\n", err)
	}

	fmt.Printf("Client Streaming Response: %v \n", res)
}
