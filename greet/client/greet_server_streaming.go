package main

import (
	"context"
	"io"
	"log"
	"proto-go/greet/proto"
)

func doGreetServerStreamingClient(c proto.GreetServiceClient) {
	log.Printf("Greet from client in Server streaming")

	req := &proto.GreetRequest{
		FirstName: "Avinash",
	}

	stream, err := c.GreetServerStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greeet Server Streaming Time %v \n", err)
	}

	for true {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Warning while reading the stream %v \n", err)
		}
		if err == io.EOF {
			log.Printf("Break as error is EOF :%v", err)
			break
		}
		log.Printf("Greet Sterver Streaming %s\n", msg.Result)
	}
}
