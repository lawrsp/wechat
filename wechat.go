package wechat

type Wechat struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func New(appid, secret string) *Wechat {
	return &Wechat{appid, secret}
}
