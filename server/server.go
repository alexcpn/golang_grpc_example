package server

import (
	"context"
	 log "github.com/sirupsen/logrus"
	 "net"
	 "fmt"
	"google.golang.org/grpc"
	pb "interfaces/test_server"
)

// server is used to implement the server
type Server struct{}

func (s *Server)Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error){
	log.Printf("Received Emailid: %v", in.EmailId)
	log.Printf("Received Query: %v", in.Query)
	
	// Note this is the only place we use validate
	err := in.Validate()
	if err != nil {
		log.Warn("SearchRequest validation failed: %v", err)
		response := pb.SearchResponse{SearchResponse: "Invalid Inputs" }
		response.Validate()
		return &response, nil
	}else {
		response := pb.SearchResponse{SearchResponse: "Some Valid response from server " }
		response.Validate()
		return &response, nil
	}
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
