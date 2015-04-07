package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/daneharrigan/nyx/context"
	"github.com/daneharrigan/nyx/nameserver"
)

const Timeout = 100 * time.Millisecond

func NewTransport(p Proxy) context.Acceptor {
	t := &transport{proxy: p}
	t.reverseProxy = &httputil.ReverseProxy{
		Director: t.Director,
		Transport: &http.Transport{
			Dial: t.Dial,
		},
	}

	return t
}

type transport struct {
	proxy        Proxy
	reverseProxy *httputil.ReverseProxy
	context      *context.Context
	err          error
}

func (t *transport) Accept(c *context.Context) {
	t.context = c
	t.reverseProxy.ServeHTTP(c.ResponseWriter, c.Request)
}

func (t *transport) Director(r *http.Request) {
	record, err := t.proxy.Nameserver().Lookup(r.Host)
	if err != nil {
		t.err = err
		return
	}

	node := record.Nodes[0] // TODO: randomly select
	r.Host = node.Host
	r.URL.Host = node.Host

	switch node.Protocol {
	case nameserver.HTTP:
		r.URL.Scheme = "http"
	case nameserver.HTTPS:
		r.URL.Scheme = "https"
	}
}

func (t *transport) Dial(network, address string) (net.Conn, error) {
	return net.DialTimeout(network, address, Timeout)
}
