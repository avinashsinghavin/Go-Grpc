package main

import (
	"fmt"
	"io"
	"proto-go/calculator/proto"
)

func (receiver Server) AverageClientStreaming(stream proto.CalculatorService_AverageClientStreamingServer) error {
	fmt.Println("Call Average Client Streaming in server module")
	res := 0
	n := 0

	for {
		no, err := stream.Recv()
		fmt.Printf("Server Average No: %v    %v\n", no, err)
		if err != nil {
			if err == io.EOF {
				if res > 0 && n > 0 {
					return stream.SendAndClose(&proto.AverageResponse{
						Total: int32(res / n),
					})
				} else {
					return stream.SendAndClose(&proto.AverageResponse{Total: 0})
				}
			} else {
				fmt.Printf("Error while reading steam recv %v \n", err)
				break
			}
		}
		n++
		res += int(no.Number)
	}

	return nil
}
