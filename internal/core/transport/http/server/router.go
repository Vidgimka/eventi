package core_http_server

import (
	"fmt"
	"net/http"
)

type ApiVertion string

var (
	ApiVertion1 = ApiVertion("v1")
	ApiVertion2 = ApiVertion("v2")
	ApiVertion3 = ApiVertion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVertion ApiVertion
}

func NewAPIVersionRouter(apiVertion ApiVertion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVertion: apiVertion,
	}
}

func (r *APIVersionRouter) RegisterRouts(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Metod, route.Path)
		r.Handle(pattern, route.Handler)
	}
}
