package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"proto-go/calculator/proto"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didi not connect with:%v\n", err)
	}

	defer conn.Close()
	c := proto.NewCalculatorServiceClient(conn)
	doCalculate(c)
}
