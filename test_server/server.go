package main

import (
	"context"
	"log"
	 "net"
	 "fmt"
	"google.golang.org/grpc"
	pb "interfaces"
)

// server is used to implement the server
type Server struct{}

func (s *Server)Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error){
	log.Printf("Received: %v", in.EmailId)
	return &pb.SearchResponse{SearchResponse: "some response from server " }, nil
}


const (
	port = ":50051"
)
func main() {
	//coding/testgo# go build cassandratest && ./cassandratest -ni 10 -nq 1000000 -del=false -pno=10
	fmt.Println("Go protobuffer test")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
