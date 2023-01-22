package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"proto-go/greet/proto"
)

var addr string = "0.0.0.5051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listion on: %v\n", err)
	}

	log.Fatalf("Listining on %s \n", addr)

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to server %v \n", err)
	}
}
