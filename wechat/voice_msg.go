package wechat

import (
	"github.com/ckeyer/alone/tuling"
)

type VoiceMsg struct {
	MsgInfo

	MediaId     string `xml:"MediaId"`
	Format      string `xml:"Format"`
	Recognition string `xml:"Recognition"`
	MsgId       int64  `xml:"MsgId"`
}

// default auto response
func (m *VoiceMsg) MsgHandle() (*ResponseMessage, error) {
	return NewTextResposeMessage(m.ToUserName, m.FromUserName, tuling.Reply(m.Recognition)), nil
}
