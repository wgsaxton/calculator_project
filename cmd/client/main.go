package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	calculatorpb "classes.com/maarek_classes/calculator_project/rpc"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is Client")

	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("Client must only take 2 integers to add together")
	}
	fmt.Println(len(args))

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorClient(cc)

	//fmt.Printf("Created client: %f", c)

	v1, _ := strconv.ParseInt(args[0], 10, 32)
	nv1 := int32(v1)

	v2, _ := strconv.ParseInt(args[1], 10, 32)
	nv2 := int32(v2)

	doUnary(c, nv1, nv2)
}

func doUnary(c calculatorpb.CalculatorClient, v1 int32, v2 int32) {
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Values{
			Val1: v1,
			Val2: v2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v\n", err)
	}
	log.Printf("Response from Sum: %v\n", res.Result)
}
