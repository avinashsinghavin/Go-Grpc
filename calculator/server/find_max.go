package main

import (
	"io"
	"log"
	"proto-go/calculator/proto"
)

func (receiver *Server) FindMaxNumberBiDirectional(server proto.CalculatorService_FindMaxNumberBiDirectionalServer) error {

	log.Println("FindMaxNumber BiDirection function call successfully")
	var maxi int64 = 0
	var greatest int64 = 0
	for true {
		req, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client stream %v\n", err)
		}
		if req.X > req.Y {
			maxi = req.X
			if greatest < req.X {
				greatest = req.X
			}
		} else {
			maxi = req.Y
			if greatest < req.GetY() {
				greatest = req.GetY()
			}
		}
		err = server.Send(&proto.MaxResponse{
			Response: maxi,
			Max:      greatest,
		})
		if err != nil {
			log.Printf("Error while sending response to client %v\n", err)
		}
	}
	return nil
}
