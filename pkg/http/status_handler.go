package http

import (
	"io"
	"net/http"
)

// boosted from @matryer
type StatusHandler int

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := int(s)
	w.WriteHeader(code)
	io.WriteString(w, http.StatusText(code))
}

var (
	NotFoundHandler       = StatusHandler(404)
	NotImplementedHandler = StatusHandler(501)
	NotLegalHandler       = StatusHandler(451)
)
