package wechat

import (
	"fmt"
	"net/http"
)

const jsCode2SessionURLFormat = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=SECRET&js_code=%s&grant_type=authorization_code"

// JSCode2Session
// doc: https://developers.weixin.qq.com/miniprogram/dev/api/api-login.html#wxloginobject
// url: https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func (w *Wechat) JSCode2SessionRequest(jsCode string) (*http.Request, error) {

	url := fmt.Sprintf(jsCode2SessionURLFormat, w.AppID, jsCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create jscode2session request failed: %v", err)
	}

	return req, nil
}
