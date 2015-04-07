package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/daneharrigan/nyx/context"
	"github.com/daneharrigan/nyx/nameserver"
)

type Transporter interface {
	context.Acceptor
}

func NewTransporter(p Proxy) Transporter {
	t := &transporter{proxy: p}
	t.reverseProxy = &httputil.ReverseProxy{
		Director:  t.Director,
		Transport: t,
	}

	return t
}

type transporter struct {
	proxy        Proxy
	reverseProxy *httputil.ReverseProxy
	context      *context.Context
	err          error
}

func (t *transporter) Accept(c *context.Context) {
	t.context = c
	t.reverseProxy.ServeHTTP(c.ResponseWriter, c.Request)
}

func (t *transporter) Director(r *http.Request) {
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
		r.URL.Scheme = "HTTPS"
	}
}

func (t *transporter) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.err != nil {
		return nil, t.err
	}

	return nil, nil
}
