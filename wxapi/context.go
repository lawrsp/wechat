package wxapi

import "context"

type wxapiKey struct{}

var wxapiKeyOfContext = wxapiKey{}

func NewContext(ctx context.Context, c *WXAPI) context.Context {
	return context.WithValue(ctx, wxapiKeyOfContext, c)
}

func FromContext(ctx context.Context) (*WXAPI, bool) {
	c, ok := ctx.Value(wxapiKeyOfContext).(*WXAPI)
	return c, ok
}
