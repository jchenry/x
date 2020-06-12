package main

import (
	"net/http"
	"text/template"
)

func render(p string, m string, body []byte, w http.ResponseWriter) (err error) {
	if tmpl, err := template.ParseFiles("page.tmpl.html"); err == nil {
		return tmpl.Execute(w, struct {
			Mode string
			Body string
			Page string
		}{m, string(body), p})
	}
	return err
}
