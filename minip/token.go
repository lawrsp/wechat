package minip

import (
	"fmt"
	"net/http"
)

const getTokenURLFormat = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// AccessTokenResponse
type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// AccessTokenRequest
func (mp *MiniP) GetAccessTokenRequest() (*http.Request, error) {
	url := fmt.Sprintf(getTokenURLFormat, mp.AppID, mp.AppSecret)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create getAccessToken request failed: %v", err)
	}

	return req, nil
}

// GetAccessToken
// https://developers.weixin.qq.com/miniprogram/dev/api/open-api/access-token/getAccessToken.html
func (mp *MiniP) GetAccessToken() (*GetAccessTokenResponse, error) {
	client := DefaultClient

	req, err := mp.GetAccessTokenRequest()
	if err != nil {
		return nil, fmt.Errorf("create getAccessToken request failed: %v", err)
	}

	out := struct {
		*Error
		*GetAccessTokenResponse
	}{
		&Error{},
		&GetAccessTokenResponse{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get getAccessToken reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.Error
	}

	return out.GetAccessTokenResponse, nil
}
