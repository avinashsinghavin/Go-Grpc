package main

import (
	"context"
	"io"
	"log"
	"proto-go/greet/proto"
	"time"
)

func greetBiDirectional(client proto.GreetServiceClient) {
	log.Printf("Greet Every one BiDirectional from Client")

	stream, err := client.GreetEveryOneBiDirectional(context.Background())
	if err != nil {
		log.Printf("Error while creating stream in client Side %v\n", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Avinash"},
		{FirstName: "Akash"},
		{FirstName: "Aman"},
		{FirstName: "Amit"},
		{FirstName: "Ankit"},
		{FirstName: "Amrit"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending Request %v \n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
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
				log.Printf("Error while receiving: %v \n", err)
				break
			}
			log.Printf("Received %v \n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
