package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s , services.NewCalculatorServer())
	
	fmt.Println("gRPC server listening on port 9090")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
