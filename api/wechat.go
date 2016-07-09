package api

import (
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/alone/wechat"
)

func AuthServerMW(ctx *RequestContext) {
	signature := ctx.params["signature"]
	timestamp := ctx.params["timestamp"]
	nonce := ctx.params["nonce"]
	echostr := ctx.params["echostr"]

	if !wechat.Auth(signature, timestamp, nonce, echostr) {
		ctx.render.Error(http.StatusUnauthorized)
		log.Warningf("auth failed, params: %v", ctx.params)
		return
	}

	if ctx.req.Method == "GET" {
		log.Info("wechat auth success")
		ctx.render.Text(200, echostr)
	} else {
		log.Debugf("wechat auth success")
	}
}

func MsgHandle(w http.ResponseWriter, req *http.Request, ctx *RequestContext) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("get body failed, error: %s", err.Error())
		ctx.render.Error(http.StatusBadRequest)
		return
	}

	data, err := wechat.MsgHandle(bs)
	if err != nil {
		ctx.render.Error(http.StatusBadRequest)
		log.Errorf("handle wechat message failed, error: %s", err.Error())
		return
	}

	ctx.render.XML(200, data)
}
