package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/vitalksocialnetwork/ProtoApi/auth/public"
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
	r, err := c.Ping(ctx, &pb.Ping{Timestamp: time.Now().Unix()})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Pong: %d", r.GetTimestamp())
}
