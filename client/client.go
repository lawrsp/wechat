package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Parser interface {
	Parse([]byte, interface{}) error
}

// Client
type Client struct {
	*http.Client
	Parser
}

// NewClient
func NewClient(p Parser) *Client {
	return &Client{&http.Client{}, p}
}

// ClientOf
func ClientOf(c *http.Client, p Parser) *Client {
	return &Client{c, p}
}

// GetReply
func (c *Client) GetReply(req *http.Request, out interface{}) error {
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %v", err)
	}

	if err := c.Parser.Parse(body, out); err != nil {
		return fmt.Errorf("parse body failed: %v", err)
	}

	return nil
}
