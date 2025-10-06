package main

import (
	"context"
	"log"

	pb "example.com/greet/proto"
)

func (*Server) 	Greet(ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
log.Printf(" Greet Received: %v", in.GetFirstName())
return &pb.GreetResponse{
	Result: "Hello " + in.GetFirstName(),	
},nil



}
