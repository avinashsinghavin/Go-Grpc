package main

import (
	"context"
	"log"
	"proto-go/calculator/proto"
)

func (receiver *Server) Calculate(ctx context.Context, in *proto.CalculationRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	switch in.GetOperation() {
	case "+":
		result := int64(in.Y + in.X)
		return &proto.CalculatorResponse{Result: result}, nil
	case "-":
		result := int64(in.GetX() - in.GetY())
		return &proto.CalculatorResponse{Result: result}, nil
	case "*":
		result := int64(in.GetX() * in.GetY())
		return &proto.CalculatorResponse{Result: result}, nil
	case "/":
		if in.GetY() < in.GetX() {
			result := int64(in.GetX() / in.GetY())
			return &proto.CalculatorResponse{Result: result}, nil
		}
		if in.GetY() > in.GetX() {
			result := int64(in.GetY() / in.GetX())
			return &proto.CalculatorResponse{Result: result}, nil
		}

	}
	return &proto.CalculatorResponse{Result: 0}, nil
}
