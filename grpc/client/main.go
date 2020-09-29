package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/syhlion/2020-12th-ithome30day-istio-loadbalance/grpc/pingpong"

	"google.golang.org/grpc"
)

func main() {
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingPongServiceClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Receive request")
		fmt.Fprintf(w, "Hello World")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rr, err := c.Ping(ctx, &pb.PingRequest{Text: "test"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Receive Grpc Response: %s", rr.GetText())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
