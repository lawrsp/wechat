package minip

import "context"

type minipKey struct{}

var minipKeyOfContext = minipKey{}

func NewContext(ctx context.Context, c *MiniP) context.Context {
	return context.WithValue(ctx, minipKeyOfContext, c)
}

func FromContext(ctx context.Context) (*MiniP, bool) {
	c, ok := ctx.Value(minipKeyOfContext).(*MiniP)
	return c, ok
}
