package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"happy-car/shared/auth/token"
	"io/ioutil"
	"os"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	publicKey *rsa.PublicKey
	verifier  tokenVerifier
}

// Interceptor creates a grpc auth interceptor.
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open publicKeyFile: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read publicKeyFile: %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse publicKey: %v", err)
	}
	i := &interceptor{
		publicKey: pubKey,
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accountID, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}
	return handler(ContextWithAccountID(ctx, AccountID(accountID)), req)
}

type accountIDKey struct{}

// AccountID defines account id object.
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// ContextWithAccountID returns a new context with the accountID.
func ContextWithAccountID(c context.Context, accountID AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, accountID)
}

func tokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]

		}
	}
	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

// AccountIDFromContext gets account id from context.
// Returns unauthenticated error if no account id is available.
func AccountIDFromContext(c context.Context) (AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
