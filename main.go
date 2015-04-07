package main

import (
	"os"

	"github.com/daneharrigan/nyx/logger"
	"github.com/daneharrigan/nyx/middleware"
	"github.com/daneharrigan/nyx/nameserver"
	"github.com/daneharrigan/nyx/proxy"
	"github.com/daneharrigan/nyx/server"
)

var log = logger.New(os.Stderr, "ns=nyx")

func main() {
	log.Println("at=start")

	ns := nameserver.New()
	ns.Add("dane", &nameserver.Node{
		Host:     "daneharrigan.com",
		Port:     80,
		Protocol: nameserver.HTTP,
	})

	p := proxy.New()
	p.SetLogger(log)
	p.SetNameserver(ns)

	s := server.New()
	s.SetLogger(log)
	s.SetProxy(p)
	s.Use(new(middleware.RequestIDHandler))

	if err := s.Listen(); err != nil {
		log.Println("fn=Listen error=%q", err)
	}

	log.Println("at=finish")
}
