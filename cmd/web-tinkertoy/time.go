package main

import (
	"io"
	"net/http"
	"time"
)

func (s *server) handleTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, time.Now().String())
	}
}
