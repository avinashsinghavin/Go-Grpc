package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"proto-go/calculator/proto"
)

var addr = "localhost:50052"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v \n\n", err)
	}

	defer lis.Close()

	log.Printf("Listing at %v\n\n", addr)

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Field to serve: %v\n", err)
	}
}
