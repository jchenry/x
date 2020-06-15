package http

import (
	"fmt"
	"net/http"
)

// MutliHandler takes a map of http methods to handlers
// and returns a handler which will run the the mapped hander
// based on a request's method
func MutliHandler(h map[string]http.Handler) (http.HandlerFunc, error) {
	m := map[string]bool{
		http.MethodHead:    true,
		http.MethodPost:    true,
		http.MethodPut:     true,
		http.MethodPatch:   true,
		http.MethodDelete:  true,
		http.MethodConnect: true,
		http.MethodOptions: true,
		http.MethodTrace:   true,
	}

	for verb := range h {
		if _, ok := m[verb]; !ok {
			return nil, fmt.Errorf("invalid HTTP method: %s", verb)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if hdlr, ok := h[r.Method]; ok {
			hdlr.ServeHTTP(w, r)
		} else {
			NotFoundHandler.ServeHTTP(w, r)
		}

	}, nil
}
