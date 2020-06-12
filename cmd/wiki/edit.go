package main

import (
	"net/http"
	"os"
)

func edit(pageName string, w http.ResponseWriter, r *http.Request) (err error) {
	if body, err := getFile(pageName, os.O_RDWR|os.O_CREATE); err == nil {
		return render(pageName, "edit", body, w)
	}
	return
}
