package wechat

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Msg struct {
	Id int64 `xml:"-",orm:"Id"`

	MsgType      string `xml:"MsgType,cdata",orm:"MsgType"`
	Event        string `xml:"Event",orm:"Event"`
	ToUserName   string `xml:"ToUserName,cdata",orm:"ToUserName"`
	FromUserName string `xml:"FromUserName,cdata",orm:"FromUserName"`
	CreateTime   int    `xml:"CreateTime",orm:"CreateTime"`

	CreatedLocal time.Time `orm:"auto_now_add;type(datetime)"`
}

// ReceiveMsg 消息的接受
func (m *Msg) ReceiveMsg() interface{} {
	switch m.MsgType {
	case "text":
		m.ReceiveTextMsg()
	case "image":
		m.ReceiveImageMsg()
	case "voice":
		m.ReceiveVoiceMsg()
	case "video":
		m.ReceiveVideoMsg()
	case "location":
		m.ReceiveLocationMsg()
	case "link":
		m.ReceiveLinkMsg()
	default:
		m.WriteText("观察君出海去啦~~~")
	}
	return nil
}

func (m *Msg) ReceiveEvent() {
	switch m.Event {
	case "subscribe":
		m.ReceiveSubscribeEvent()
	case "unsubscribe":
		m.ReceiveUnsubscribeEvent()
	case "SCAN":
		m.ReceiveScanEvent()
	case "LOCATION":
		m.ReceiveLocationEvent()
	case "CLICK":
		m.ReceiveClickEvent()
	case "VIEW":
		m.ReceiveViewEvent()
	}
}

func (m *Msg) WriteText(data string) {
	xmlfmt := `<xml><ToUserName><![CDATA[%s]]>CD</ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>`
	body := fmt.Sprintf(xmlfmt, m.FromUserName, m.ToUserName, time.Now().Unix(), data)
	log.Info("receive body: ", string(m.content))
	log.Info("send body: ", body)
	m.w.WriteHeader(http.StatusOK)
	m.w.Write([]byte(body))
}
