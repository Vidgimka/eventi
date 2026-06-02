package core_http_middleware

import (
	"net/http"

	"github.com/google/uuid"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
