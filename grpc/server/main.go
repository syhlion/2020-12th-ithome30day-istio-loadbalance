package main

import (
	"context"
	"log"
	"net"

	pb "github.com/syhlion/2020-12th-ithome30day-istio-loadbalance/grpc/pingpong"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingPongServiceServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	log.Printf("Receive Grpc Request: %v", in.GetText())
	return &pb.PongResponse{Text: "Hello " + in.GetText()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingPongServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
