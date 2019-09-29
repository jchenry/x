package http

import (
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type Middleware interface {
	UseHandler(handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Router interface {
	ServeFiles(path string, root http.FileSystem)
	AddHandler(method, path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Service interface {
	Register(uriBase string, restServer *Server)
}

type ServiceFunc func(uriBase string, restServer *Server)

func (f ServiceFunc) Register(uriBase string, restServer *Server) {
	f(uriBase, restServer)
}

var docString = "%s  \t%s\t- %s"

type Server struct {
	router     *httprouter.Router
	middleware Middleware
}

func NewServer(m Middleware) *Server {
	s := &Server{
		router:     httprouter.New(),
		middleware: m,
	}

	s.middleware.UseHandler(s.router)

	return s
}

func (r *Server) GET(path string, documentation string, handle http.Handler) *Server {
	r.handle("GET", path, documentation, handle)
	return r
}
func (r *Server) PATCH(path string, documentation string, handle http.Handler) *Server {
	r.handle("PATCH", path, documentation, handle)

	return r
}
func (r *Server) POST(path string, documentation string, handle http.Handler) *Server {
	r.handle("POST", path, documentation, handle)

	return r
}
func (r *Server) PUT(path string, documentation string, handle http.Handler) *Server {
	r.handle("PUT", path, documentation, handle)

	return r
}
func (r *Server) DELETE(path string, documentation string, handle http.Handler) *Server {
	r.handle("DELETE", path, documentation, handle)

	return r
}
func (r *Server) handle(method, path string, documentation string, handler http.Handler) {
	log.Printf(docString, method, path, documentation)
	r.router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		if req.Form == nil {
			req.Form = url.Values{}
		}
		for _, param := range params { // stuffing values back into request.Form to honor the handler contract
			req.Form.Add(param.Key, param.Value)
		}
		handler.ServeHTTP(w, req)
	})
}

func (r *Server) Banner(banner string) *Server {
	log.Printf(banner)
	return r
}

func (r *Server) Service(basePath string, service Service) *Server {
	service.Register(basePath, r)
	return r
}
func (r *Server) Static(path string, root http.FileSystem) *Server {
	r.router.ServeFiles(path, root)
	return r
}
func (r *Server) Middleware(handler http.Handler) *Server {
	r.middleware.UseHandler(handler)
	return r
}

func (r *Server) Run(addr string) {
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r.middleware))
}

func (r *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.middleware.ServeHTTP(w, req)
}
