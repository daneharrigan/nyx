package server

import (
	"net/http"

	"github.com/daneharrigan/nyx/context"
	"github.com/daneharrigan/nyx/logger"
	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/proxy"
)

type Server interface {
	SetLogger(logger.Logger)
	SetProxy(proxy.Proxy)
	Use(middleware.Middleware)
	Logger() logger.Logger
	Listen() error
}

func New() Server {
	return new(server)
}

type server struct {
	proxy      proxy.Proxy
	logger     logger.Logger
	middleware []middleware.Middleware
	http       *http.Server
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
	srv := &http.Server{Addr: ":8080", Handler: s}
	return srv.ListenAndServe()
}

func (s *server) Logger() logger.Logger {
	return s.logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := context.New(w, r)

	var acceptor context.Acceptor = s.proxy
	for _, m := range s.middleware {
		acceptor = m.Build(acceptor)
	}

	acceptor.Accept(c)
}
