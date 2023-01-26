package main

import (
	"fmt"
	"log"
	"proto-go/greet/proto"
)

func (s *Server) GreetServerStreaming(in *proto.GreetRequest, stream proto.GreetService_GreetServerStreamingServer) error {
	log.Printf("Greet Server Streaming (Many time Gret message from Server) %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s number %d", in.FirstName, i)

		stream.Send(&proto.GreetResponse{Result: res})
	}

	return nil
}