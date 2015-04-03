package proxy

import (
	"net/http"

	"github.com/daneharrigan/nyx/nameserver"
	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/logger"
)

type Proxy interface {
	http.Handler
	SetLogger(logger.Logger)
	SetNameserver(nameserver.Nameserver)
	Use(middleware.Middleware)
}

func New() Proxy {
	return new(proxy)
}

type proxy struct {
	logger logger.Logger
	ns nameserver.Nameserver
	middleware []middleware.Middleware
}

func (p *proxy) SetLogger(l logger.Logger) {
	p.logger = l
}

func (p *proxy) SetNameserver(ns nameserver.Nameserver) {
	p.ns = ns
}

func (p *proxy) Use(m middleware.Middleware) {
	p.middleware = append(p.middleware, m)
}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = p.reverseProxy
	for _, m := range p.middleware {
		handler = m.Build(handler)
	}

	p.logger.Printf("at=proxy fn=ServeHTTP request_id=%s",
		r.Header.Get("Request-Id"))
}
