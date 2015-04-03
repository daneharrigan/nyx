package server

import (
	"net/http"

	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/proxy"
	"github.com/daneharrigan/nyx/logger"
)

type Server interface {
	SetLogger(logger.Logger)
	SetProxy(proxy.Proxy)
	Use(middleware.Middleware)
	Listen() error
}

func New() Server {
	return new(server)
}

type server struct {
	proxy proxy.Proxy
	logger logger.Logger
	middleware []middleware.Middleware
	http *http.Server
}

func (s *server) SetLogger(l logger.Logger) {
	s.logger = l
}

func (s *server) SetProxy(p proxy.Proxy) {
	s.proxy = p
}

func (s *server) Use(m middleware.Middleware) {
	s.middleware = append(s.middleware, m)
}

func (s *server) Listen() error {
	var handler http.Handler = s.proxy
	for _, m := range s.middleware {
		handler = m.Build(handler)
	}

	s.http = &http.Server{Addr: ":8080", Handler: handler}
	return s.http.ListenAndServe()
}
