package main

import (
	"context"
	"io"
	"log"

	pb "example.com/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient){
	
	stream, err := c.GreetManyTimes(context.Background(),&pb.GreetRequest{
		FirstName: "John",
	})
	if err != nil {
		panic(err)
	}

	for {
msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}

	
}