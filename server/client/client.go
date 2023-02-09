package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	tsClient := trippb.NewTripServiceClient(conn)
	response, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "123456",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v", err)
	}
	fmt.Println(response)
}
