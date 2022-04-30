package trip

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rentalpb "happy-car/rental/api/gen/v1"
	"happy-car/shared/auth"
)

// Service implements a trip service.
type Service struct {
	Logger *zap.Logger
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	// get accountID from context
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip", zap.String("start", request.Start), zap.String("account_id", aid.String()))
	return nil, status.Error(codes.Unimplemented, "")
}
