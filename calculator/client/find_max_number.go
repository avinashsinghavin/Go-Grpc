package main

import (
	"context"
	"io"
	"log"
	"proto-go/calculator/proto"
	"time"
)

func findMaxNumber(client proto.CalculatorServiceClient) {
	var list = []int64{1, 5, 3, 6, 2, 20}
	var i = 0
	req, err := client.FindMaxNumberBiDirectional(context.Background())
	if err != nil {
		log.Printf("Error while caling Server Function FindMaxNumberBiDirectional %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for i < len(list)-1 {
			req.Send(&proto.NumbersRequest{X: list[i], Y: list[i+1]})
			time.Sleep(1 * time.Second)
			i += 1
		}
		req.CloseSend()
	}()

	go func() {
		for true {
			res, err := req.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while reading data from server side %v\n", err)
				break
			}
			log.Printf("Largert Number is %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}
