package minip

import (
	"encoding/json"
	"fmt"

	"github.com/lawrsp/wechat/client"
)

type MiniP struct {
	AppID     string
	AppSecret string
}

func New(appid, secret string) *MiniP {
	return &MiniP{appid, secret}
}

type parser struct{}

func (*parser) Parse(bs []byte, out interface{}) error {
	if err := json.Unmarshal(bs, out); err != nil {
		return fmt.Errorf("parse failed: %v", err)
	}

	return nil
}

var DefaultClient = client.NewClient(&parser{})
