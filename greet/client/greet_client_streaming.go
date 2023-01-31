package main

import (
	"context"
	"log"
	"proto-go/greet/proto"
	"time"
)

func greetClientStreaming(client proto.GreetServiceClient) {
	log.Println("long greet from client was invoked")
	reqs := []*proto.GreetRequest{
		{
			FirstName: "Avinash",
		}, {
			FirstName: "Pritam",
		}, {
			FirstName: "Sourav",
		},
	}

	stream, err := client.GreetClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Long Greet or client streaming %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while receiving response from client Streaming %v\n", err)
	}

	log.Printf("Client Streaming Response: %v \n", res)
}
