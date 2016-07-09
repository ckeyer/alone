package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/ckeyer/alone/api"
	"github.com/ckeyer/alone/tuling"
	"github.com/ckeyer/alone/wechat"
	"github.com/ckeyer/commons/config"
	"github.com/spf13/viper"
)

func init() {
	config.Init("alone")
}

func main() {
	wechat.Init()
	tuling.Init()
	logrus.Debug("...")
	logrus.Info(viper.GetString("logging.level"))
	api.Serve(viper.GetString("app.listening"))
}
