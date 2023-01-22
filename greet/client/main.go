package client

import (
	"google.golang.org/grpc"
	"log"
	"proto-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Failes to connect %v\n", err)
	}

	defer conn.Close()

	c := proto.NewGreetServiceClient(conn)

	doGreet(c)
}
