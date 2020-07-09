package rest

import (
	gohttp "net/http"

	"github.com/jchenry/x/net/http"
)

// EntityHandler returns a handler that provides restful verbs, following a CRUD model
func EntityHandler(
	get gohttp.Handler,
	post gohttp.Handler,
	put gohttp.Handler,
	delete gohttp.Handler,
) gohttp.HandlerFunc {
	h, _ := http.MutliHandler(map[string]gohttp.Handler{
		gohttp.MethodGet:    get,
		gohttp.MethodPost:   post,
		gohttp.MethodPut:    put,
		gohttp.MethodDelete: delete,
	})
	return h
}
