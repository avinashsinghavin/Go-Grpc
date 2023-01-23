package main

import (
	"context"
	"log"
	"proto-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Geet function was invoked with %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
