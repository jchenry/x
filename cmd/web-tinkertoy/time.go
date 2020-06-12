package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jchenry/jchenry/arvelie"
	"github.com/jchenry/jchenry/neralie"
)

func (s *server) handleTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(200)
		io.WriteString(w, t.String())
		io.WriteString(w, fmt.Sprintf("\n%s %s",
			arvelie.FromDate(t),
			neralie.FromTime(t)))
	}
}
