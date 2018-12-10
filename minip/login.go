package minip

import (
	"fmt"
	"net/http"
)

const code2AccessTokenURLFormat = "https://api.weixin.qq.com/sns/code2accessToken?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

// Code2AccessTokenResponse
type Code2AccessTokenResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
}

// Code2AccessTokenRequest
func (mp *MiniP) Code2AccessTokenRequest(jsCode string) (*http.Request, error) {
	url := fmt.Sprintf(code2AccessTokenURLFormat, mp.AppID, mp.AppSecret, jsCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create code2accessToken request failed: %v", err)
	}

	return req, nil
}

// Code2AccessToken
// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/code2accessToken.html
func (mp *MiniP) Code2AccessToken(jsCode string) (*Code2AccessTokenResponse, error) {
	client := DefaultClient

	req, err := mp.Code2AccessTokenRequest(jsCode)
	if err != nil {
		return nil, fmt.Errorf("create code2accessToken request failed: %v", err)
	}

	out := struct {
		*Error
		*Code2AccessTokenResponse
	}{
		&Error{},
		&Code2AccessTokenResponse{},
	}
	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get code2accessToken reply failed: %v", err)
	}

	if out.Error.IsError() {
		return nil, out.Error
	}

	return out.Code2AccessTokenResponse, nil
}
