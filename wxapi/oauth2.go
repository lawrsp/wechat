package wxapi

import (
	"fmt"
	"net/http"
)

var oauth2AcessTokenURLFormat = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"

type OAuth2AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (mp *WXAPI) OAuth2AccessTokenRequest(code string) (*http.Request, error) {
	url := fmt.Sprintf(oauth2AcessTokenURLFormat, mp.AppID, mp.AppSecret, code)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create oauth2accessToken request failed: %v", err)
	}

	return req, nil
}

func (mp *WXAPI) OAuth2AccessToken(code string) (*OAuth2AccessTokenResponse, error) {
	client := DefaultClient

	req, err := mp.OAuth2AccessTokenRequest(code)
	if err != nil {
		return nil, fmt.Errorf("create oauth2accessToken request failed: %v", err)
	}

	out := struct {
		*Error
		*OAuth2AccessTokenResponse
	}{
		&Error{},
		&OAuth2AccessTokenResponse{},
	}
	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get oauth2accessToken reply failed: %v", err)
	}

	if out.Error.IsError() {
		return nil, out.Error
	}

	return out.OAuth2AccessTokenResponse, nil
}
