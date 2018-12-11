package wxapi

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
func (mp *WXAPI) GetAccessTokenRequest() (*http.Request, error) {
	url := fmt.Sprintf(getTokenURLFormat, mp.AppID, mp.AppSecret)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create getAccessToken request failed: %v", err)
	}

	return req, nil
}

// GetAccessToken
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
func (mp *WXAPI) GetAccessToken() (*GetAccessTokenResponse, error) {
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

var getCallbackIPURLFormat = "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s"

type GetCallbackIPResponse struct {
	IPList []string `json:"ip_list"`
}

func (*WXAPI) GetCallbackIPRequest(accessToken string) (*http.Request, error) {
	url := fmt.Sprintf(getCallbackIPURLFormat, accessToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create getCallbackIP request failed: %v", err)
	}

	return req, nil

}

// GetCallbackIP
// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140187
func (mp *WXAPI) GetCallbackIP(accessToken string) (*GetCallbackIPResponse, error) {
	client := DefaultClient

	req, err := mp.GetCallbackIPRequest(accessToken)
	if err != nil {
		return nil, fmt.Errorf("create getCallbackIP request failed: %v", err)
	}

	out := struct {
		*Error
		*GetCallbackIPResponse
	}{
		&Error{},
		&GetCallbackIPResponse{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get getCallbackIP reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.Error
	}

	return out.GetCallbackIPResponse, nil
}
