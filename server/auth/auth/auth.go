package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"happy-car/auth/api/gen/v1"
	"happy-car/auth/auth/dao"
)

type Service struct {
	OpenIdResolver OpenIdResolver
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}

// OpenIdResolver 接口应该由使用者定义
// resolves an authorization code to an openid.
type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	//s.Logger.Info("received code", zap.String("code", request.Code))
	openID, err := s.OpenIdResolver.Resolve(request.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(ctx, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: "token for accountID " + accountID,
		ExpiresIn:   7200,
	}, nil
}
