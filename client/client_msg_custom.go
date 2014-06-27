package client

import (
	"errors"
	"github.com/chanxuehong/wechat/message/custom"
)

// 发送客服消息功能都一样, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgCustomSend(msg interface{}) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := messageCustomSendURL(token)

	var result Error
	if err = c.postJSON(_url, msg, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return &result
	}
	return
}

// 发送客服消息, 文本.
func (c *Client) MsgCustomSendText(msg *custom.Text) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图片.
func (c *Client) MsgCustomSendImage(msg *custom.Image) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 语音.
func (c *Client) MsgCustomSendVoice(msg *custom.Voice) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 视频.
func (c *Client) MsgCustomSendVideo(msg *custom.Video) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 音乐.
func (c *Client) MsgCustomSendMusic(msg *custom.Music) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图文.
func (c *Client) MsgCustomSendNews(msg *custom.News) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}