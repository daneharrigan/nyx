package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/daneharrigan/nyx/context"
)

const Timeout = 3 * time.Millisecond

type Handler struct {
	proxy Proxy
	reverseProxy *httputil.ReverseProxy
	context *context.Context
	context.Acceptor
}

func NewHandler(p Proxy) *Handler {
	handler := &Handler{proxy: p}
	handler.reverseProxy = &httputil.ReverseProxy{
		Transport: &http.Transport{
			Dial: handler.Dial
		},
		Director: handler.Director,
	}

	return handler
}

func (h *Handler) Accept(c *context.Context) {
	h.context = c
	h.reverseProxy.ServeHTTP(c.ResponseWriter, c.Request)
}

func (h *Handler) Dial(network, address string) (net.Conn, error) {
	return net.DialTimeout(network, address, Timeout)
}

func (h *Handler) Director(r *http.Request) {
	// look up destination
	// decorate *http.Request
}
