package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dkalthoff/grpc-customers/customerspb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Get(context context.Context, req *customerspb.GetQuery) (*customerspb.Customer, error) {
	fmt.Println("Got a new Get Customer request")
	customerID := req.GetCustomerId()

	result := &customerspb.Customer{CustomerId: customerID, FirstName: "John", LastName: "Doe"}
	return result, nil
}

func main() {
	fmt.Println("Starting Customers server on port 1971")
	lis, err := net.Listen("tcp", "0.0.0.0:1971")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	t := grpc.NewServer()
	customerspb.RegisterCustomersServer(t, &server{})
	if err := t.Serve(lis); err != nil {
		log.Fatalf("Error while serving customers : %v", err)
	} else {
		fmt.Println("serving customers")
	}
}
