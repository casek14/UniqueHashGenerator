package main

import (
	"github.com/casek14/UniqueHashGenerator/hash"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	port     = ":50051"
	etcdUrl  = "127.0.0.1"
	etcdPort = "2379"
)

func main() {
	url, err := os.LookupEnv("DBURL")
	if err {
		url = etcdUrl
	}
	dbPort, err := os.LookupEnv("DBPORT")
	if err {
		dbPort = etcdPort
	}
	lis, errs := net.Listen("tcp", port)
	if errs != nil {
		log.Fatalf("FAILED TO LISTEN: %v", errs)
	}
	log.Printf("Starting GRPC HASH SERVER ON PORT %s", port)
	hServer := hash.NewHasServer(url, dbPort)
	log.Printf("SERVER CONFIG %s", hServer.Client.Config.Endpoints)
	s := grpc.NewServer()
	hash.RegisterHashServer(s, hServer)
	s.Serve(lis)
}
