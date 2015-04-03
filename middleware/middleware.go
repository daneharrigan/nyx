package middleware

import (
	"net/http"
	"code.google.com/p/go-uuid/uuid"
)

type Middleware interface {
	http.Handler
	Build(http.Handler) http.Handler
}

type RequestIDHandler struct {
	Handler http.Handler
}

func (h *RequestIDHandler) Build(handler http.Handler) http.Handler {
	h.Handler = handler
	return h
}

func (h *RequestIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Request-Id") == "" {
		r.Header.Set("Request-Id", uuid.New())
	}

	w.Header().Set("Request-Id", r.Header.Get("Request-Id"))
	h.Handler.ServeHTTP(w, r)
}
