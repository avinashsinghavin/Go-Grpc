package main

import (
	"log"
	"proto-go/calculator/proto"
)

func (receiver *Server) PrimeDecomposition(in *proto.NumberRequest, stream proto.CalculatorService_PrimeDecompositionServer) error {
	log.Printf("Server layer operation to find prime decomposition number %v\n", in)
	k := 2
	n := 210
	for n > 1 {
		if n%k == 0 {
			stream.Send(&proto.PrimeNoResponse{PrimeNo: int64(k)})
			n = n / k
		} else {
			k = k + 1
		}

	}
	return nil
}
