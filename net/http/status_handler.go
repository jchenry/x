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
	_, _ = io.WriteString(w, http.StatusText(code))
}

var (
	NotFoundHandler       = StatusHandler(http.StatusNotFound)
	NotImplementedHandler = StatusHandler(http.StatusNotImplemented)
	NotLegalHandler       = StatusHandler(http.StatusUnavailableForLegalReasons)
	NotAllowedHandler     = StatusHandler(http.StatusMethodNotAllowed)
)
