package wxapi

import (
	"encoding/json"
	"fmt"

	"github.com/lawrsp/wechat/client"
)

type WXAPI struct {
	AppID     string
	AppSecret string
}

func New(appid, secret string) *WXAPI {
	return &WXAPI{appid, secret}
}

type parser struct{}

func (*parser) Parse(bs []byte, out interface{}) error {
	if err := json.Unmarshal(bs, out); err != nil {
		return fmt.Errorf("parse failed: %v", err)
	}

	return nil
}

var DefaultClient = client.NewClient(&parser{})

func (wx *WXAPI) accessToken() string {
	return ""
}
