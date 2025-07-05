package grpc

import (
	"context"

	"credit-card-validator/internal/service"
	pb "credit-card-validator/pkg/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCardValidatorServer
	validator *service.Validator
	logger    *logrus.Logger
}

func NewServer(validator *service.Validator, logger *logrus.Logger) *Server {
	return &Server{
		validator: validator,
		logger:    logger,
	}
}

func (s *Server) RegisterServer(grpcServer *grpc.Server) {
	pb.RegisterCardValidatorServer(grpcServer, s)
}

func (s *Server) ValidateCard(ctx context.Context, req *pb.ValidateCardRequest) (*pb.ValidateCardResponse, error) {
	s.logger.WithField("request_id", ctx.Value("request_id")).Info("gRPC ValidateCard called")

	result, err := s.validator.ValidateCard(ctx, req.CardNumber)
	if err != nil {
		s.logger.WithError(err).Error("Card validation failed")
		return nil, status.Errorf(codes.Internal, "validation failed: %v", err)
	}

	res := &pb.ValidateCardResponse{
		Valid:      result.Valid,
		CardType:   string(result.CardType),
		CardNumber: result.CardNumber,
		Scheme:     result.Scheme,
		CardBrand:  result.CardBrand,
		CardKind:   result.CardKind,
	}

	// Add country if available
	if result.Country.Name != "" {
		res.Country = &pb.Country{
			Name:      result.Country.Name,
			Alpha2:    result.Country.Alpha2,
			Currency:  result.Country.Currency,
			Emoji:     result.Country.Emoji,
			Latitude:  int32(result.Country.Latitude),
			Longitude: int32(result.Country.Longitude),
		}
	}

	// Add bank if available
	if result.Bank.Name != "" {
		res.Bank = &pb.Bank{
			Name:  result.Bank.Name,
			Url:   result.Bank.URL,
			Phone: result.Bank.Phone,
		}
	}

	return res, nil
}
