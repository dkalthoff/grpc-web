package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dkalthoff/grpc-customers/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(context context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Println("Got a new Add request")
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	result := &calculatorpb.AddResponse{Result: sum}
	return result, nil
}

func main() {
	fmt.Println("Starting Calculator server on port 1961")
	lis, err := net.Listen("tcp", "0.0.0.0:1961")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving calculator : %v", err)
	} else {
		fmt.Println("serving calculator")
	}
}
