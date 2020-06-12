package http

import (
	"log"
	"net/http"
	go_http "net/http"
)

type Middleware interface {
	go_http.Handler
	UseHandler(handler http.Handler)
}

type Router interface {
	go_http.Handler
	ServeFiles(path string, root go_http.FileSystem)
	AddHandler(method, path string, handler go_http.Handler)
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
	router     Router
	middleware Middleware
}

func NewServer(m Middleware, r Router) *Server {
	s := &Server{
		router:     r,
		middleware: m,
	}

	s.middleware.UseHandler(s.router)

	return s
}

func (r *Server) Get(path string, documentation string, handle go_http.Handler) *Server {
	r.handle("GET", path, documentation, handle)
	return r
}
func (r *Server) Patch(path string, documentation string, handle go_http.Handler) *Server {
	r.handle("PATCH", path, documentation, handle)

	return r
}
func (r *Server) Post(path string, documentation string, handle go_http.Handler) *Server {
	r.handle("POST", path, documentation, handle)

	return r
}
func (r *Server) Put(path string, documentation string, handle go_http.Handler) *Server {
	r.handle("PUT", path, documentation, handle)

	return r
}
func (r *Server) Delete(path string, documentation string, handle go_http.Handler) *Server {
	r.handle("DELETE", path, documentation, handle)

	return r
}
func (r *Server) handle(method, path string, documentation string, handler go_http.Handler) {
	log.Printf(docString, method, path, documentation)
	r.router.AddHandler(method, path, handler)
}

func (r *Server) Banner(banner string) *Server {
	log.Printf(banner)
	return r
}

func (r *Server) Service(basePath string, service Service) *Server {
	service.Register(basePath, r)
	return r
}
func (r *Server) Static(path string, root go_http.FileSystem) *Server {
	r.router.ServeFiles(path, root)
	return r
}
func (r *Server) Middleware(handler go_http.Handler) *Server {
	r.middleware.UseHandler(handler)
	return r
}

func (r *Server) Run(addr string) {
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r.middleware))
}

func (r *Server) ServeHTTP(w go_http.ResponseWriter, req *go_http.Request) {
	r.middleware.ServeHTTP(w, req)
}
