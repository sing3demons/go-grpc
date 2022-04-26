package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(creds))
	defer cc.Close()
	if err != nil {
		log.Fatal(err)
	}
	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("sing")
	if err != nil {
		log.Fatal(err)
	}
}
