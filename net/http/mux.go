// based on the "way" package by matt ryer
// Copyright (c) 2016 Mat Ryer

package http

import (
	"context"
	"net/http"
	"strings"
)

// contextKey is the context key type for storing
// parameters in context.Context.
type contextKey string

// Router routes HTTP requests.
type ServeMux struct {
	routes []*route
	// NotFound is the http.Handler to call when no routes
	// match. By default uses http.NotFoundHandler().
	NotFound http.Handler
}

// NewRouter makes a new Router.
func NewServeMux() *ServeMux {
	return &ServeMux{
		NotFound: http.NotFoundHandler(),
	}
}

func (r *ServeMux) pathSegments(p string) []string {
	return strings.Split(strings.Trim(p, "/"), "/")
}

// Handle adds a handler with the specified  pattern.
// Pattern can contain path segments such as: /item/:id which is
// accessible via the Param function.
// If pattern ends with trailing /, it acts as a prefix.
func (r *ServeMux) Handle(pattern string, handler http.Handler) {
	route := &route{
		segs:    r.pathSegments(pattern),
		handler: handler,
		prefix:  strings.HasSuffix(pattern, "/") || strings.HasSuffix(pattern, "..."),
	}
	r.routes = append(r.routes, route)
}

// HandleFunc is the http.HandlerFunc alternative to http.Handle.
func (r *ServeMux) HandleFunc(pattern string, fn http.HandlerFunc) {
	r.Handle(pattern, fn)
}

// ServeHTTP routes the incoming http.Request based on path
func (r *ServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	segs := r.pathSegments(req.URL.Path)
	for _, route := range r.routes {
		if ctx, ok := route.match(req.Context(), r, segs); ok {
			route.handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
	}
	r.NotFound.ServeHTTP(w, req)
}

// Param gets the path parameter from the specified Context.
// Returns an empty string if the parameter was not found.
func Param(ctx context.Context, param string) string {
	vStr, ok := ctx.Value(contextKey(param)).(string)
	if !ok {
		return ""
	}
	return vStr
}

type route struct {
	segs    []string
	handler http.Handler
	prefix  bool
}

func (r *route) match(ctx context.Context, router *ServeMux, segs []string) (context.Context, bool) {
	if len(segs) > len(r.segs) && !r.prefix {
		return nil, false
	}
	for i, seg := range r.segs {
		if i > len(segs)-1 {
			return nil, false
		}
		isParam := false
		if strings.HasPrefix(seg, "{") {
			isParam = true
			seg = strings.Trim(seg, "{}")
		}
		if !isParam { // verbatim check
			if strings.HasSuffix(seg, "...") {
				if strings.HasPrefix(segs[i], seg[:len(seg)-3]) {
					return ctx, true
				}
			}
			if seg != segs[i] {
				return nil, false
			}
		}
		if isParam {
			ctx = context.WithValue(ctx, contextKey(seg), segs[i])
		}
	}
	return ctx, true
}
