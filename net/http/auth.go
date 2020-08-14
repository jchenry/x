package http

import (
	"net/http"
	"os"
)

func BasicAuth(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !(user == os.Getenv("WIKI_USERNAME") && pass == os.Getenv("WIKI_PASSWORD")) {
			w.Header().Set("WWW-Authenticate", `Basic realm="wiki"`)
			http.Error(w, "Unauthorized.", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}
