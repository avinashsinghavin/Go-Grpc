package main

import (
	"context"
	"log"
	"proto-go/greet/proto"
)

func doGreet(c proto.GreetServiceClient) {
	log.Printf("do greet was invoked")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "Avinash",
	})

	if err != nil {
		log.Fatalf("could not greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
