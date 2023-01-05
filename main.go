package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	server "github.com/gopiesy/grpc-health-server/health-server"
	health "github.com/gopiesy/grpc-health-server/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC health server should listen for TCP conn.")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	health.RegisterHealthServer(grpcServer, server.NewHealthServer())
	log.Printf("Initializing gRPC health server on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
