package wepay

import "context"

type wepayKey struct{}

var wepayKeyOfContext = wepayKey{}

func NewContext(ctx context.Context, c *WePay) context.Context {
	return context.WithValue(ctx, wepayKeyOfContext, c)
}

func FromContext(ctx context.Context) (*WePay, bool) {
	c, ok := ctx.Value(wepayKeyOfContext).(*WePay)
	return c, ok
}
