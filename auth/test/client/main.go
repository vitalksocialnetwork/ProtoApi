package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vitalksocialnetwork/ProtoApi/auth/gen-go"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{timestamp: time.Second})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Ping: %s", r.GetMessage())
}
