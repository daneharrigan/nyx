package proxy

import (
	"github.com/daneharrigan/nyx/context"
	"github.com/daneharrigan/nyx/logger"
	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/nameserver"
)

type Proxy interface {
	context.Acceptor
	SetLogger(logger.Logger)
	SetNameserver(nameserver.Nameserver)
	Logger() logger.Logger
	Nameserver() nameserver.Nameserver
	Use(middleware.Middleware)
}

func New() Proxy {
	return new(proxy)
}

type proxy struct {
	logger     logger.Logger
	nameserver nameserver.Nameserver
	middleware []middleware.Middleware
}

func (p *proxy) SetLogger(l logger.Logger) {
	p.logger = l
}

func (p *proxy) SetNameserver(ns nameserver.Nameserver) {
	p.nameserver = ns
}

func (p *proxy) Use(m middleware.Middleware) {
	p.middleware = append(p.middleware, m)
}

func (p *proxy) Logger() logger.Logger {
	return p.logger
}

func (p *proxy) Nameserver() nameserver.Nameserver {
	return p.nameserver
}

func (p *proxy) Accept(c *context.Context) {
	var acceptor context.Acceptor = NewTransport(p)
	for _, m := range p.middleware {
		acceptor = m.Build(acceptor)
	}

	acceptor.Accept(c)
}
