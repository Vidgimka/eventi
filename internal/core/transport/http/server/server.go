package core_http_server

import "net/http"

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
}

func NewHTTPServer(config Config) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
	}
}

func (h *HTTPServer) Run() error {
	server := http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}
}
