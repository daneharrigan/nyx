package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/daneharrigan/nyx/nameserver"
	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/logger"
	"github.com/daneharrigan/nyx/context"
)

type Proxy interface {
	context.Acceptor
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

func (p *proxy) Accept(c *context.Context) {
	var acceptor context.acceptor = NewHandler(p)
	for _, m := range p.middleware {
		acceptor = m.Build(acceptor)
	}

	acceptor.Accept(c)
}
