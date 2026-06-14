package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Vidgimka/eventi/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    core_logger.Logger
}

func NewHTTPServer(config Config, log core_logger.Logger) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
		log:    log,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVertion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.log.Warn("start HTTP server", zap.String("add", h.config.Addr))

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shotdown HTTP server...")

		shotdownContext, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shotdownContext); err != nil {
			_ = server.Close()
			return fmt.Errorf("shotdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")
	}
	return nil
}
