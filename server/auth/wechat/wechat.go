package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v3"
)



type Service struct {
	AppID     string
	AppSecret string
}



func (s *Service) Resolve(code string) (string, error) {
	cli := weapp.NewClient(s.AppID, s.AppSecret)
	resp, err := cli.Login(code)
	if err != nil {
		return "", fmt.Errorf("wechat Login: %v", err)
	}
	if err := resp.GetResponseError(); err != nil {
		return "", fmt.Errorf("wechat get Response; %v", err)
	}
	return resp.OpenID, nil
}