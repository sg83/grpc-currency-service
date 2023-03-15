package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	pb "github.com/sg83/grpc-currency-service/proto"
)

// Define a struct that implements the CurrencyServiceServer interface
type currencyServer struct {
	log hclog.Logger
	pb.UnimplementedCurrencyServiceServer
}

func NewCurrencyServer(l hclog.Logger) *currencyServer {
	return &currencyServer{l, pb.UnimplementedCurrencyServiceServer{}}
}

// Implement the GetExchangeRate method
func (s *currencyServer) GetExchangeRate(ctx context.Context, req *pb.ExchangeRateRequest) (*pb.ExchangeRateResponse, error) {
	// In a real implementation, you would retrieve the exchange rate data from some external API or database
	// For this example, we'll just return a hardcoded exchange rate of 1.23
	exchangeRate := 1.23
	s.log.Info("Handle GetExchangeRate", "base:", req.BaseCurrency, "destination:", req.TargetCurrency)
	// Create and return an ExchangeRateResponse message containing the exchange rate
	return &pb.ExchangeRateResponse{
		ExchangeRate: exchangeRate,
	}, nil
}
