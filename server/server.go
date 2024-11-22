package server
import (
	"context"
	"fmt"
	"log"
	"net"
	pb"github.com/kasyap1234/protobuf-example/gen/person"
	"google.golang.org/grpc"
)
type greeterServer struct {
	pb.UnimplementedGreeterServer
}


func (s *greeterServer) SayHello(ctx context.Context, req *pb.Person)(*pb.Person,error){
	log.Printf("Received: %v", req.GetName())
	return &pb.Person{Name: "Hello " + req.GetName()}, nil

}
func StartServer(){
	server :=grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{})
	list,err :=net.Listen("tcp",":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %f", err)
	}
	fmt.Println("Server started on port 50051")

}