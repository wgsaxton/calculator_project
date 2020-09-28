package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorpb "classes.com/maarek_classes/calculator_project/rpc"

	"google.golang.org/grpc"
)

// Server See if this works

// Sum - see if caps works
func Sum(ctx context.Context, requ *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calling server Sum func\n")
	val1 := requ.GetSum().GetVal1()
	val2 := requ.GetSum().GetVal2()

	result := val1 + val2
	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := &calculatorpb.CalculatorService{
		Sum: Sum,
	}
	calculatorpb.RegisterCalculatorService(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
