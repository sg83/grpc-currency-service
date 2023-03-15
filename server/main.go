package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/sg83/grpc-currency-service/server/proto"
)

// Define a struct that implements the CurrencyServiceServer interface
type currencyServer struct{}

// Implement the GetExchangeRate method
func (s *currencyServer) GetExchangeRate(ctx context.Context, req *pb.ExchangeRateRequest) (*pb.ExchangeRateResponse, error) {
	// In a real implementation, you would retrieve the exchange rate data from some external API or database
	// For this example, we'll just return a hardcoded exchange rate of 1.23
	exchangeRate := 1.23

	// Create and return an ExchangeRateResponse message containing the exchange rate
	return &pb.ExchangeRateResponse{
		ExchangeRate: exchangeRate,
	}, nil
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the currencyServer as the implementation of the CurrencyServiceServer interface
	pb.RegisterCurrencyServiceServer(server, &currencyServer{})

	// Listen on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	log.Printf("Starting server on port %v", listener.Addr().String())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
