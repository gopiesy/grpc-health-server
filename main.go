package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	server "github.com/gopiesy/grpc-server/policy-server"
	"github.com/gopiesy/project-protos/policies"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC server should listen TCP conn.")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	policies.RegisterPolicyServiceServer(grpcServer, server.PolicyServiceServer{})
	log.Printf("Initializing gRPC server on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
