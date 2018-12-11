package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140837

var createTagURLFormat = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"

type createTagInfo struct {
	Name string `json:"name"`
}
type createTagReq struct {
	Tag createTagInfo `json:"tag"`
}

// CreateTagRequest
func (wx *WXAPI) CreateTagRequest(name string) (*http.Request, error) {
	token := wx.accessToken()

	url := fmt.Sprintf(createTagURLFormat, token)

	info := &createTagReq{createTagInfo{Name: name}}

	bs, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	return req, nil
}

// TagInfo
type TagInfo struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// CreateTagResponse
type CreateTagResponse struct {
	Tag TagInfo `json:"tag"`
}

// CreateTag
func (wx *WXAPI) CreateTag(name string) (*CreateTagResponse, error) {

	client := DefaultClient
	req, err := wx.CreateTagRequest(name)
	if err != nil {
		return nil, fmt.Errorf("createTag failed: %v", err)
	}

	out := struct {
		*Error
		*CreateTagResponse
	}{
		&Error{},
		&CreateTagResponse{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get createTag reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.Error
	}

	return out.CreateTagResponse, nil
}

var getTagsURLFormat = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"

// GetTagsRequest
func (wx *WXAPI) GetTagsRequest() (*http.Request, error) {
	token := wx.accessToken()
	url := fmt.Sprintf(getTagsURLFormat, token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// GetTagsResponse
type GetTagsResponse struct {
	Tags []*TagInfo `json:"tags"`
}

// GetTags
func (wx *WXAPI) GetTags() (*GetTagsResponse, error) {

	client := DefaultClient
	req, err := wx.GetTagsRequest()
	if err != nil {
		return nil, fmt.Errorf("create getTags request failed: %v", err)
	}

	out := struct {
		*Error
		*GetTagsResponse
	}{
		&Error{},
		&GetTagsResponse{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get getTags reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.Error
	}

	return out.GetTagsResponse, nil
}
