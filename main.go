package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"

	server "github.com/gopiesy/grpc-health-server/health-server"
	health "github.com/gopiesy/grpc-health-server/proto"
	"github.pie.apple.com/isd-dev/go-applecerts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC health server should listen for TCP conn.")
	cert = flag.String("cert", "/Users/gopichandkomma/Desktop/certs/previewhealthchecker.edge.apple.com.chain.pem", "server cert")
	key  = flag.String("key", "/Users/gopichandkomma/Desktop/certs/previewhealthchecker.edge.apple.com.chain.pem", "server key")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// options
	var opts []grpc.ServerOption
	if cert != nil && key != nil {
		cert, err := tls.LoadX509KeyPair(*cert, *key)
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig, err := applecerts.TLSConfig()
		if err != nil {
			log.Fatal(err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	grpcServer := grpc.NewServer(opts...)
	health.RegisterHealthServer(grpcServer, server.NewHealthServer())
	log.Printf("Initializing gRPC health server on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
