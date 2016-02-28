package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/ckeyer/alone/api"
	"github.com/ckeyer/alone/conf"
	logpkg "github.com/ckeyer/go-log"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	API_PREFIX = ""
)

var (
	log = logpkg.GetDefaultLogger("wechat")

	configPath = flag.String("c", "conf/v1.json", "configure file path")
)

func main() {
	config, err := conf.LoadConfigFile(*configPath)
	if err != nil {
		log.Fatal("load config err, ", err)
	}
	addr := ":" + strconv.Itoa(config.App.Port)
	m := NewMartini()
	m.Group(API_PREFIX, api.APIRoute())

	log.Notice("server listening at: ", addr)
	if err := http.ListenAndServe(addr, m); err != nil {
		log.Fatal(err)
	}
}

func NewMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{Martini: m, Router: r}
}
