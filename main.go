package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	server "github.com/gopiesy/grpc-health-server/health-server"
	health "github.com/gopiesy/grpc-health-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 9111, "Port on which gRPC health server should listen for TCP conn.")
	root = flag.String("root", "./certs/cacert.pem", "root CA")
	cert = flag.String("cert", "./certs/server.pem", "server cert")
	key  = flag.String("key", "./certs/server.key", "server key")
)

func main() {
	flag.Parse()

	// Load the server certificate and its key
	serverCert, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		log.Fatalf("Failed to load server certificate and key. %s.", err)
	}

	// Load the CA certificate
	trustedCert, err := os.ReadFile(*root)
	if err != nil {
		log.Fatalf("Failed to load trusted certificate. %s.", err)
	}

	// Put the CA certificate to certificate pool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		log.Fatalf("Failed to append trusted certificate to certificate pool. %s.", err)
	}

	// Create the TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// options
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(cred))

	grpcServer := grpc.NewServer(opts...)
	health.RegisterHealthServer(grpcServer, server.NewHealthServer())
	log.Printf("Initializing gRPC health server on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
