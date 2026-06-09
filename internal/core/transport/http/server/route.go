package core_http_server

import "net/http"

type Route struct {
	Metod   string
	Path    string
	Handler http.HandlerFunc
}

func NewRote(metod, path string, handler http.HandlerFunc) Route {
	return Route{
		Metod:   metod,
		Path:    path,
		Handler: handler,
	}
}
