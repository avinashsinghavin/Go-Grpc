package main

import (
	"fmt"
	"io"
	"log"
	"proto-go/greet/proto"
)

func (s *Server) GreetClientStreaming(stream proto.GreetService_GreetClientStreamingServer) error {
	fmt.Println("Greet from client streaming and server was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&proto.GreetResponse{
					Result: res,
				})
			} else {
				log.Fatalf("Error while reading client Stream %s \n", err)
				break
			}
		}
		res += fmt.Sprintf("Hello %s !\n", req.FirstName)
	}
	return nil
}
