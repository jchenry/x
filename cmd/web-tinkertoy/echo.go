package main

import (
	"net/http"
	"sync"
	"text/template"
)

func (s *server) handleEcho() http.HandlerFunc {
	var (
		init   sync.Once
		tmpl   *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tmpl, tplerr = template.ParseFiles("echoform.tmpl")
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			s.echoMessage = r.Form.Get("msg")
			http.Redirect(w, r, "/echo", 301)
			return
		case http.MethodGet:
			if err := tmpl.Execute(w, map[string]string{"Message": s.echoMessage}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
