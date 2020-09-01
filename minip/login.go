package minip

import (
	"fmt"
	"net/http"
)

const code2SessionURLFormat = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

// Code2SessionResponse
type Code2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
}

// Code2SessionRequest
func (mp *MiniP) Code2SessionRequest(jsCode string) (*http.Request, error) {
	url := fmt.Sprintf(code2SessionURLFormat, mp.AppID, mp.AppSecret, jsCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create code2accessToken request failed: %v", err)
	}

	return req, nil
}

// Code2Session
// https://developers.weixin.qq.com/miniprogram/dev/api/code2Session.html
func (mp *MiniP) Code2Session(jsCode string) (*Code2SessionResponse, error) {
	client := DefaultClient

	req, err := mp.Code2SessionRequest(jsCode)
	if err != nil {
		return nil, fmt.Errorf("create code2accessToken request failed: %v", err)
	}

	out := struct {
		*Error
		*Code2SessionResponse
	}{
		&Error{},
		&Code2SessionResponse{},
	}
	if err := client.GetReply(req, &out); err != nil {
		return nil, fmt.Errorf("get code2accessToken reply failed: %v", err)
	}

	if out.Error.IsError() {
		return nil, out.Error
	}

	return out.Code2SessionResponse, nil
}
