package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// 验证token并解析accountID

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies a token and returns accountID
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	// 验证token的过程
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not StandardClaims")
	}
	// 验claim（验证过期时间之类的）
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid: %v", err)
	}

	// 返回accountID
	return clm.Subject, nil
}
