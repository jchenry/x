package http

import (
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type JulienschmidtHTTPRouter struct {
	httprouter.Router
}

func NewJulienschmidtHTTPRouter() *JulienschmidtHTTPRouter {
	return &JulienschmidtHTTPRouter{
		httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
		},
	}
}

func (j *JulienschmidtHTTPRouter) AddHandler(method, path string, handler http.Handler) {
	j.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		if req.Form == nil {
			req.Form = url.Values{}
		}
		for _, param := range params {
			// stuffing values back into request.Form to honor the handler contract
			req.Form.Add(param.Key, param.Value)
		}
		handler.ServeHTTP(w, req)
	})
}
