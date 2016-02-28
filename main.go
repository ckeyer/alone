package main

import (
	"flag"
	"strconv"

	"github.com/ckeyer/alone/api"
	"github.com/ckeyer/alone/conf"
	logpkg "github.com/ckeyer/go-log"
)

const (
	API_PREFIX = ""
)

var (
	log = logpkg.GetDefaultLogger("wechat")

	configPath = flag.String("c", "conf/v1.json", "configure file path")
)

func main() {
	flag.Parse()
	config, err := conf.LoadConfigFile(*configPath)
	if err != nil {
		log.Fatal("load config err, ", err)
	}
	addr := ":" + strconv.Itoa(config.App.Port)
	api.Serve(addr)
}
