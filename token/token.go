package token

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lawrsp/wechat/client"
	"github.com/lawrsp/wechat/errors"
)

type parser struct{}

func (*parser) Parse(bs []byte, out interface{}) error {
	if err := json.Unmarshal(bs, out); err != nil {
		return fmt.Errorf("parse failed: %v", err)
	}

	return nil
}

var DefaultClient = client.NewClient(&parser{})

const getTokenURLFormat = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// TokenInfo
type TokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// GetAccessTokenRequest
func GetAccessTokenRequest(appID, appSecret string) (*http.Request, error) {
	url := fmt.Sprintf(getTokenURLFormat, appID, appSecret)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create getAccessToken request failed: %v", err)
	}

	return req, nil
}

// GetAccessToken
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func GetAccessToken(appID, appSecret string) (*TokenInfo, error) {
	client := DefaultClient

	req, err := GetAccessTokenRequest(appID, appSecret)
	if err != nil {
		return nil, fmt.Errorf("create getAccessToken request failed: %v", err)
	}

	out := struct {
		*errors.WxError
		*TokenInfo
	}{
		&errors.WxError{},
		&TokenInfo{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get getAccessToken reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.WxError
	}

	return out.TokenInfo, nil
}
