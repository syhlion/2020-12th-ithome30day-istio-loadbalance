package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/syhlion/2020-12th-ithome30day-istio-loadbalance/pingpong"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingPongServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c.Ping(ctx, &pb.PingRequest{Text: "test"})
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s")
}
