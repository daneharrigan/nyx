package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	err := http.ListenAndServe(":8080", new(Server))
	if err != nil {
		log.Printf("fn=ListenAndServe error=%q", err)
	}
}

type Server struct {}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("http://daneharrigan.com")
	if err != nil {
		log.Fatalf("fn=Parse error=%q", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = &ProxyTransport{http.DefaultTransport}
	proxy.ServeHTTP(w, r)
}

type ProxyTransport struct {
	http.RoundTripper
}

func (p *ProxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Host = r.URL.Host

	log.Printf("URL: %s\n", r.URL)
	log.Printf("Host: %s\n", r.Host)
	for k, v := range r.Header {
		log.Printf("%s: %s\n", k, v)
	}

	return p.RoundTripper.RoundTrip(r)
}
