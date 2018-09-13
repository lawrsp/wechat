package wechat

import "context"

type configKey struct{}

var wechatConfigKeyOfContext = configKey{}

func NewContext(ctx context.Context, c *Wechat) context.Context {
	return context.WithValue(ctx, wechatConfigKeyOfContext, c)
}

func FromContext(ctx context.Context) (*Wechat, bool) {
	c, ok := ctx.Value(wechatConfigKeyOfContext).(*Wechat)
	return c, ok
}
