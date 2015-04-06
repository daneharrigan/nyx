package context

import (
	"net/http"

	"code.google.com/p/go-uuid/uuid"
)

func New(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		InternalID:     uuid.New(),
	}
}

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	RequestID      string
	InternalID     string
}

type Acceptor interface {
	Accept(*Context)
}
