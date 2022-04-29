package wechat

import (
	"fmt"
	"github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppId     string
	AppSecret string
}

// Resolve 从微信后台换取openId
func (s *Service) Resolve(code string) (string, error) {
	loginResponse, err := weapp.Login(s.AppId, s.AppSecret, code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login: %v", err)
	}
	if err := loginResponse.GetResponseError(); err != nil {
		return "", fmt.Errorf("weapp response error: %v", err)
	}
	return loginResponse.OpenID, nil
}
