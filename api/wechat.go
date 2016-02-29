package api

import (
	"io/ioutil"
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
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err)
		ctx.render.Error(http.StatusBadRequest)
		return
	}

	data, err := wechat.MsgHandle(bs)
	if err != nil {
		ctx.render.Error(http.StatusBadRequest)
		log.Error(err)
		return
	}

	ctx.render.XML(http.StatusOK, data)

	log.Debug("msg type: ", msg.MsgType)
	log.Debug("msg from: ", msg.FromUserName)
	log.Debug("msg to: ", msg.ToUserName)
	log.Debug("msg create time: ", msg.CreateTime)
}
