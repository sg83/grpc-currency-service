package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-hclog"
	pb "github.com/sg83/grpc-currency-service/proto"
	"github.com/sg83/grpc-currency-service/server"
)

func main() {
	// Create a new gRPC server
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)

	//Testing purpose - Disable later
	reflection.Register(gs)

	// Register the currencyServer as the implementation of the CurrencyServiceServer interface
	pb.RegisterCurrencyServiceServer(gs, cs)

	// Listen on port 50051
	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Failed to listen: %v", err)
	}

	// Start the gRPC server
	log.Info("Starting server on port %v", listener.Addr().String())
	if err := gs.Serve(listener); err != nil {
		log.Error("Failed to serve: %v", err)
	}
}
