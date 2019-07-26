package main

import (
	"context"
	"fmt"
	"github.com/casek14/UniqueHashGenerator/hash"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func GetHash(client hash.HashClient) {
	r, err := client.GetHash(context.Background(), &hash.HashRequest{})
	if err != nil {
		log.Printf("Unable to get HASH from gRPC server: %s\n", err)
	}

	fmt.Printf("Received hash is: %s\n", r.Hash)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Cannot connect to gRPC server: %s\n", err)
	}
	defer conn.Close()
	client := hash.NewHashClient(conn)

	GetHash(client)
}
