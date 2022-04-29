package auth

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"happy-car/auth/api/gen/v1"
	"happy-car/auth/auth/dao"
	"time"
)

type Service struct {
	OpenIdResolver OpenIdResolver
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         *zap.Logger
}

// OpenIdResolver 接口应该由使用者定义
// resolves an authorization code to an openid.
type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

// TokenGenerator generates a token for the specified account.
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
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

	token, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		//AccessToken: "token for accountID " + accountID,
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
