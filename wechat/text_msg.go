package wechat

import (
	"github.com/ckeyer/alone/tuling"
)

type TextMsg struct {
	MsgInfo
	Content string `xml:"Content"`
	MsgId   int64  `xml:"MsgId"`
}

// default auto response
func (m *TextMsg) MsgHandle() (*ResponseMessage, error) {
	return NewTextResposeMessage(m.ToUserName, m.FromUserName, tuling.Reply(m.Content)), nil
}
