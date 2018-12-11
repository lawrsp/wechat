package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1458044813

var addKFAccountURLFormat = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"

type addKFAccountRequest struct {
	KFAccount string `json:"kf_account"`
	NickName  string `json:"nickname"`
	Password  string `json:"password"`
}

func (wx *WXAPI) makeKFAccountRequest(urlFormat string, account, nickname, password string) (*http.Request, error) {
	token := wx.accessToken()
	url := fmt.Sprintf(urlFormat, token)
	body := &addKFAccountRequest{
		KFAccount: account,
		NickName:  nickname,
		Password:  password,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	return req, nil

}

func (wx *WXAPI) AddKFAccountRequest(account, nickname, password string) (*http.Request, error) {
	req, err := wx.makeKFAccountRequest(addKFAccountURLFormat, account, nickname, password)
	if err != nil {
		return nil, fmt.Errorf("create addKFAccount request failed: %v", err)
	}

	return req, nil
}

// AddKFAccount
func (wx *WXAPI) AddKFAccount(account, nickname, password string) error {
	client := DefaultClient
	req, err := wx.AddKFAccountRequest(account, nickname, password)
	if err != nil {
		return fmt.Errorf("create addKFAccount request failed: %v", err)
	}

	out := struct {
		*Error
	}{
		&Error{},
	}

	if err := client.GetReply(req, out); err != nil {
		return fmt.Errorf("get addKFAccount reply failed: %v", err)
	}

	if out.IsError() {
		return out.Error
	}

	return nil

}

var updateKFAccountURLFormat = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"

// UpdateKFAccountRequest
func (wx *WXAPI) UpdateKFAccountRequest(account, nickname, password string) (*http.Request, error) {
	req, err := wx.makeKFAccountRequest(updateKFAccountURLFormat, account, nickname, password)
	if err != nil {
		return nil, fmt.Errorf("create updateKFAccount request failed: %v", err)
	}

	return req, nil
}

// UpdateKFAccount
func (wx *WXAPI) UpdateKFAccount(account, nickname, password string) error {
	client := DefaultClient
	req, err := wx.UpdateKFAccountRequest(account, nickname, password)
	if err != nil {
		return fmt.Errorf("create updateKFAccount request failed: %v", err)
	}

	out := struct {
		*Error
	}{
		&Error{},
	}

	if err := client.GetReply(req, out); err != nil {
		return fmt.Errorf("get updateKFAccount reply failed: %v", err)
	}

	if out.IsError() {
		return out.Error
	}

	return nil
}

var deleteKFAccountURLFormat = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"

func (wx *WXAPI) DeleteKFAccountRequest(account, nickname, password string) (*http.Request, error) {
	req, err := wx.makeKFAccountRequest(deleteKFAccountURLFormat, account, nickname, password)
	if err != nil {
		return nil, fmt.Errorf("create deleteKFAccount request failed: %v", err)
	}

	return req, nil
}

// DeleteKFAccount
func (wx *WXAPI) DeleteKFAccount(account, nickname, password string) error {
	client := DefaultClient
	req, err := wx.DeleteKFAccountRequest(account, nickname, password)
	if err != nil {
		return fmt.Errorf("create deleteKFAccount request failed: %v", err)
	}

	out := struct {
		*Error
	}{
		&Error{},
	}

	if err := client.GetReply(req, out); err != nil {
		return fmt.Errorf("get deleteKFAccount reply failed: %v", err)
	}

	if out.IsError() {
		return out.Error
	}

	return nil
}

// http请求方式: POST/FORM
// http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
// 调用示例：使用curl命令，用FORM表单方式上传一个多媒体文件，curl命令的具体用法请自行了解

var uploadKFHeadImageURLFormat = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s"

func (wx *WXAPI) UploadKFHeadImageReqeust(account string, filename string, image io.Reader) (*http.Request, error) {

	token := wx.accessToken()
	url := fmt.Sprintf(uploadKFHeadImageURLFormat, token, account)

	buf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buf)
	defer bodyWriter.Close()
	fileWriter, err := bodyWriter.CreateFormFile("media", filename)
	if err != nil {
		return nil, fmt.Errorf("create uploadKFHeadImage failed: %v", err)
	}

	_, err = io.Copy(fileWriter, image)
	if err != nil {
		return nil, fmt.Errorf("create uploadKFHeadImage failed: %v", err)
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func (wx *WXAPI) UploadKFHeadImage(account, filename string, image io.Reader) error {
	client := DefaultClient
	req, err := wx.UploadKFHeadImageReqeust(account, filename, image)
	if err != nil {
		return err
	}

	out := struct {
		*Error
	}{
		&Error{},
	}

	if err := client.GetReply(req, out); err != nil {
		return fmt.Errorf("get UploadKFHeadImage reply failed: %v", err)
	}

	if out.IsError() {
		return out.Error
	}

	return nil
}

var getKFListURLFormat = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"

type KFInfo struct {
	KFAccount      string `json:"kf_account"`
	KFNick         string `json:"kf_nick"`
	KFID           string `json:"kf_id"`
	KFHeadImageURL string `json:"kf_headimgurl"`
}
type GetKFListResponse struct {
	KFList []*KFInfo `json:"kf_list"`
}

func (wx *WXAPI) GetKFListRequest() (*http.Request, error) {
	token := wx.accessToken()
	url := fmt.Sprintf(getKFListURLFormat, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (wx *WXAPI) GetKFList() (*GetKFListResponse, error) {

	client := DefaultClient
	req, err := wx.GetKFListRequest()
	if err != nil {
		return nil, fmt.Errorf("create getKFList request failed: %v", err)
	}

	out := struct {
		*Error
		*GetKFListResponse
	}{
		&Error{},
		&GetKFListResponse{},
	}

	if err := client.GetReply(req, out); err != nil {
		return nil, fmt.Errorf("get getKFList reply failed: %v", err)
	}

	if out.IsError() {
		return nil, out.Error
	}

	return out.GetKFListResponse, nil
}

type MessageType string

const (
	MessageTypeText            MessageType = "text"            //文本
	MessageTypeImage           MessageType = "image"           //图片
	MessageTypeVoice           MessageType = "voice"           //语音
	MessageTypeVideo           MessageType = "video"           //视频消息
	MessageTypeMusic           MessageType = "music"           //音乐消息
	MessageTypeNews            MessageType = "news"            //图文消息(点击跳转到外链)
	MessageTypeMPNews          MessageType = "mpnews"          //图文消息(点击跳转到图文消息页面)
	MessageTypeWXCard          MessageType = "wxcard"          //卡券
	MessageTypeMiniProgramPage MessageType = "miniprogrampage" //小程序
)

// Message
type Message interface {
	messageType() MessageType
}

// TextMessage
type TextMessage struct {
	Content string `json:"content"`
}

// ImageMessage
type ImageMessage struct {
	MediaID string `json:"media_id"`
}

// VoiceMessage
type VoiceMessage struct {
	MediaID string `json:"media_id"`
}

// VideoMessage
type VideoMessage struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// MusicMessage
type MusicMessage struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicURL     string `json:"musicurl"`
	HQMusicURL   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// ArticleMsg
type ArticleMsg struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// NewsMessage
type NewsMessage struct {
	Articles []*ArticleMsg `json:"articles"`
}

// MPNewsMessage
type MPNewsMessage struct {
	MediaID string `json:"media_id"`
}

// WXCardMessage
type WXCardMessage struct {
	CardID string `json:"card_id"`
}

// MiniProgramPageMessage
type MiniProgramPageMessage struct {
	Title        string `json:"title"`
	AppID        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

func (*TextMessage) messageType() MessageType {
	return MessageTypeText
}
func (*ImageMessage) messageType() MessageType {
	return MessageTypeImage
}
func (*VoiceMessage) messageType() MessageType {
	return MessageTypeVoice
}
func (*VideoMessage) messageType() MessageType {
	return MessageTypeVideo
}
func (*MusicMessage) messageType() MessageType {
	return MessageTypeMusic
}
func (*NewsMessage) messageType() MessageType {
	return MessageTypeNews
}
func (*MPNewsMessage) messageType() MessageType {
	return MessageTypeMPNews
}
func (*WXCardMessage) messageType() MessageType {
	return MessageTypeWXCard
}
func (*MiniProgramPageMessage) messageType() MessageType {
	return MessageTypeMiniProgramPage
}

type CustomService struct {
	KFAccount string `json:"kf_account"`
}

var sendMessageURLFormat = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

// SendKFMessageRequest
func (wx *WXAPI) SendKFMessageRequest(msg Message, toUser string, kfAccount string) (*http.Request, error) {
	accessToken := wx.accessToken()
	url := fmt.Sprintf(sendMessageURLFormat, accessToken)

	msgType := string(msg.messageType())
	body := map[string]interface{}{
		"touser":  toUser,
		"msgtype": msgType,
		msgType:   msg,
	}
	if kfAccount != "" {
		body["customservice"] = CustomService{kfAccount}
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("create sendMessage request failed: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("create getCallbackIP request failed: %v", err)
	}

	return req, nil
}

// SendKFMessage
func (mp *WXAPI) SendKFMessage(message Message, toUser string, kfAccount string) error {

	client := DefaultClient
	req, err := mp.SendKFMessageRequest(message, toUser, kfAccount)
	if err != nil {
		return fmt.Errorf("sendMessage failed: %v", err)
	}

	out := struct {
		*Error
	}{
		&Error{},
	}

	if err := client.GetReply(req, out); err != nil {
		return fmt.Errorf("get sendMessage reply failed: %v", err)
	}

	if out.IsError() {
		return out.Error
	}

	return nil
}
