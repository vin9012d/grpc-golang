package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "example.com/greet/proto"
)
var addr string = "localhost:50051"

func main() {
conn,err := grpc.NewClient(addr,grpc.WithTransportCredentials(insecure.NewCredentials()));
 
if(err != nil){
	log.Fatalf("Failed to connect: %v", err)
}

c := pb.NewGreetServiceClient(conn)
doGreet(c);
doGreetManyTimes(c);
defer conn.Close()


}
