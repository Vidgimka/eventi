package user_transport_http

import (
	"net/http"

	core_http_server "github.com/Vidgimka/eventi/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
}

func NewUserHTTPHandler(userService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Metod:   http.MethodGet,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
