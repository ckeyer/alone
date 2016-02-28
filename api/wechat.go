package api

import (
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
