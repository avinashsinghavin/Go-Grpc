package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"proto-go/greet/proto"
)

var addr string = "localhost:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &Server{})

	//defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
