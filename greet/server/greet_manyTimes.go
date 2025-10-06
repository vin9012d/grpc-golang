package main

import (
	"log"
	"strconv"

	pb "example.com/greet/proto"
	"google.golang.org/grpc"
)

func (*Server) 	GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error{
log.Printf(" Greet Received: %v", in.GetFirstName())

for i:=0;i<10;i++{
	result := "Hello " + in.GetFirstName() + " number " + strconv.Itoa(i)
	stream.Send(&pb.GreetResponse{
		Result: result,
	})
	


}
return nil;
}
