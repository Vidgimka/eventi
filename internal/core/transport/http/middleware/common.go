package core_http_middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/Vidgimka/eventi/internal/core/logger"
	core_http_response "github.com/Vidgimka/eventi/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("call midleware RequestID")
			requestId := r.Header.Get(requestIDHeader)
			if requestId == "" {
				requestId = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestId)
			w.Header().Set(requestIDHeader, requestId)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("call midleware Logger")
			requestId := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("requestId", requestId),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("call midleware Panic")
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			hendlerResponse := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					hendlerResponse.PanicResponse(p, "during handle http request got unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)

		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("call midleware Trace")
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			beforeTime := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", beforeTime.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(beforeTime)),
			)
		})
	}
}
