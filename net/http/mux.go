package http

import (
	"net/http"
	"strconv"
	"strings"
)

type ServeMux struct {
	routes []route
}

func (mux *ServeMux) Handle(pattern string, handler http.Handler, pathParams ...any) {
	mux.routes = append(mux.routes, newRoute(pattern, handler, pathParams...))
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request), pathParams ...any) {
	mux.routes = append(mux.routes, newRoute(pattern, http.HandlerFunc(handler), pathParams...))
}

func (mux *ServeMux) Handler(r *http.Request) (h http.Handler, pattern string) {
	for _, rte := range mux.routes {
		switch {
		case rte.matcher(r):
			return rte.handler, rte.pattern
		}
	}
	return http.HandlerFunc(http.NotFound), ""
}

func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

type route struct {
	pattern string
	matcher func(r *http.Request) bool
	handler http.HandlerFunc
}

func newRoute(pattern string, handler http.Handler, vars ...interface{}) route {
	return route{
		pattern,
		func(r *http.Request) bool {
			return match(r.URL.Path, pattern, vars...)
		},
		handler.ServeHTTP,
	}
}

// match reports whether path matches the given pattern, which is a
// path with '+' wildcards wherever you want to use a parameter. Path
// parameters are assigned to the pointers in vars (len(vars) must be
// the number of wildcards), which must be of type *string or *int.
func match(path, pattern string, vars ...interface{}) bool {
	for ; pattern != "" && path != ""; pattern = pattern[1:] {
		switch pattern[0] {
		case '+':
			// '+' matches till next slash in path
			slash := strings.IndexByte(path, '/')
			if slash < 0 {
				slash = len(path)
			}
			segment := path[:slash]
			path = path[slash:]
			switch p := vars[0].(type) {
			case *string:
				*p = segment
			case *int:
				n, err := strconv.Atoi(segment)
				if err != nil || n < 0 {
					return false
				}
				*p = n
			default:
				panic("vars must be *string or *int")
			}
			vars = vars[1:]
		case path[0]:
			// non-'+' pattern byte must match path byte
			path = path[1:]
		default:
			return false
		}
	}
	return path == "" && pattern == ""
}
