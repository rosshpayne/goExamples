// +build OMIT

// Package gorilla provides a go.net/context.Context implementation whose Value
// method returns the values associated with a specific HTTP request in the
// github.com/gorilla/context package.
package gorilla

import (
	"net/http"
        "context"
)

// NewContext returns a Context whose Value method returns values associated
// with req using the Gorilla context package:
// http://www.gorillatoolkit.org/pkg/context
func NewContext(parent context.Context, req *http.Request) context.Context {
	return &wrapper{parent, req}
}

type wrapper struct {
	context.Context
	req *http.Request
}

type key int

const reqKey key = 0

// Value returns Gorilla's context package's value for this Context's request
// and key. It delegates to the parent Context if there is no such value.
func (ctx *wrapper) Value(key_ interface{} ) interface{} {
        if v,_ := key_.(key); v==reqKey { 
                                          return ctx.req 
        }
	return ctx.Context.Value(key_)
}

// HTTPRequest returns the *http.Request associated with ctx using NewContext,
// if any.
func HTTPRequest(ctx context.Context) (*http.Request, bool) {
	// We cannot use ctx.(*wrapper).req to get the request because ctx may
	// be a Context derived from a *wrapper. Instead, we use Value to
	// access the request if it is anywhere up the Context tree.
	req, ok := ctx.Value(reqKey).(*http.Request)
	return req, ok
}
