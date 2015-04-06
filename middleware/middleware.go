package middleware

import "github.com/daneharrigan/nyx/context"

type Middleware interface {
	context.Acceptor
	Build(context.Acceptor) context.Acceptor
}

type RequestIDHandler struct {
	Acceptor context.Acceptor
}

func (h *RequestIDHandler) Build(acceptor context.Acceptor) context.Acceptor {
	h.Acceptor = acceptor
	return h
}

func (h *RequestIDHandler) Accept(c *context.Context) {
	requestID := c.Request.Header.Get("Request-Id")
	if requestID != "" {
		c.RequestID = requestID
	} else {
		c.RequestID = c.InternalID
		c.Request.Header.Set("Request-Id", c.RequestID)
	}

	c.ResponseWriter.Header().Set("Request-Id", c.RequestID)
	h.Acceptor.Accept(c)
}
