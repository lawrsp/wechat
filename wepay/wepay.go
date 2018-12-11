package wepay

import (
	"encoding/xml"
	"fmt"

	"github.com/lawrsp/wechat/client"
)

type WePay struct {
	AppID string
	MchID string
	Key   string
}

func New(appid, mchID, key string) *WePay {
	return &WePay{appid, mchID, key}
}

type parser struct{}

func (*parser) Parse(bs []byte, out interface{}) error {
	if err := xml.Unmarshal(bs, out); err != nil {
		return fmt.Errorf("parse failed: %v", err)
	}

	return nil
}

var DefaultClient = client.NewClient(&parser{})
