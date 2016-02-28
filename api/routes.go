package api

import (
	"net/http"

	"github.com/go-martini/martini"
)

type routeHandle func(martini.Router)

func APIRoute() routeHandle {
	return func(r martini.Router) {
		r.Get("/hello", Hello)
	}
}

func Hello(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello"))
}
