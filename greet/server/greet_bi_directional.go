package main

import (
	"io"
	"log"
	"proto-go/greet/proto"
)

func (s *Server) GreetEveryOneBiDirectional(stream proto.GreetService_GreetEveryOneBiDirectionalServer) error {
	log.Println("Greet Every one was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + " !"
		log.Printf("Server received %v\n", req.FirstName)
		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Printf("Error while sending data to client %v\n", err)
		}

	}
}
