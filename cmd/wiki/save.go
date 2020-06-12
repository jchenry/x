package main

import (
	"fmt"
	"net/http"
)

func save(pageName string, w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	if err = saveFile(pageName, []byte(r.Form.Get("Text"))); err == nil {
		http.Redirect(w, r, fmt.Sprintf("/wiki/%s", pageName), http.StatusTemporaryRedirect)
	}
	return
}
