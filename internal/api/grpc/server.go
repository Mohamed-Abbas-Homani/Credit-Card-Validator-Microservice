package grpc

import (
	"context"

	"credit-card-validator/internal/service"
	pb "credit-card-validator/pkg/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	result := s.validator.ValidateCard(req.CardNumber)

	return &pb.ValidateCardResponse{
		Valid:      result.Valid,
		CardType:   string(result.CardType),
		CardNumber: result.CardNumber,
		Scheme:     result.Scheme,
		CardBrand:  result.CardBrand,
		CardKind:   result.CardKind,
		Country: &pb.Country{
			Name:      result.Country.Name,
			Alpha2:    result.Country.Alpha2,
			Currency:  result.Country.Currency,
			Emoji:     result.Country.Emoji,
			Latitude:  result.Country.Latitude,
			Longitude: result.Country.Longitude,
		},
		Bank: &pb.Bank{
			Name:  result.Bank.Name,
			Url:   result.Bank.URL,
			Phone: result.Bank.Phone,
		},
	}, nil
}
