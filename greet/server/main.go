package main

import (
	"fmt"
	"net"

	pb "example.com/greet/proto"
	"google.golang.org/grpc"
)


var addr string = "0.0.0.0:50051";

type Server struct{
pb.GreetServiceServer
}

func main() {
 fmt.Println("Starting gRPC server..."+ addr)
	lis, err := net.Listen("tcp", addr);
	
	if err != nil {
		 fmt.Println("Failed to listen:", err)
	}

	println("Server listening at", addr)
	s:= grpc.NewServer();
	pb.RegisterGreetServiceServer(s, &Server{});


	if err = s.Serve(lis); err != nil {
		 fmt.Println("Failed to serve:", err)
	}

	

	

}