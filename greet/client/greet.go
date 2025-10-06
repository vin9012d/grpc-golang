package main

import (
	"context"

	pb "example.com/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	
	res, err := c.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "John",
	})
	if err != nil {
		panic(err)
	}
	println("Response from Greet: ",res.Result)
}