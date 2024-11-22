package main 

import (
	"context"
	"fmt"
	"log"
	"time"
	pb"github.com/kasyap1234/protobuf-example/gen/person"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	person := &pb.Person{
		Name:  "Virat",
		Id:    1,
		Age:   34,
		Email: "virat@gmail.com",
	}

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
response,err := client.SayHello(ctx,person); 
if err !=nil {
	log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", response.GetName())
cancel()

}