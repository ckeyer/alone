package api

import (
	"encoding/xml"
	"net/http"

	"github.com/ckeyer/alone/wechat"
	"github.com/go-martini/martini"
)

func AuthServer(ctx *RequestContext) {
	signature := ctx.params["signature"]
	timestamp := ctx.params["timestamp"]
	nonce := ctx.params["nonce"]
	echostr := ctx.params["echostr"]

	if !wechat.Auth(signature, timestamp, nonce, echostr) {
		ctx.render.Error(http.StatusUnauthorized)
		log.Debug("auth failed")
	}
}

func MsgHandle(w http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	msg := &wechat.Msg{}
	err := xml.NewDecoder(ctx.req).Decode(msg)
	if err != nil {
		log.Error(err.Error())
		ctx.render.Error(http.StatusBadRequest)
		return
	}

	log.Debug("msg type: ", msg.MsgType)
	log.Debug("msg from: ", msg.FromUserName)
	log.Debug("msg to: ", msg.ToUserName)
	log.Debug("msg create time: ", msg.CreateTime)

}
