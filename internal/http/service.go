package http

// import "net/http"

// type Service interface {
// 	Register(s Server)
// }

// type Mux interface {
// 	Head(pattern string, handler http.Handler)
// 	Post(pattern string, handler http.Handler)
// 	Put(pattern string, handler http.Handler)
// 	Patch(pattern string, handler http.Handler)
// 	Delete(pattern string, handler http.Handler)
// 	Connect(pattern string, handler http.Handler)
// 	Options(pattern string, handler http.Handler)
// 	Trace(pattern string, handler http.Handler)
// }

// MethodGet     = "GET"
//     MethodHead    = "HEAD"
//     MethodPost    = "POST"
//     MethodPut     = "PUT"
//     MethodPatch   = "PATCH" // RFC 5789
//     MethodDelete  = "DELETE"
//     MethodConnect = "CONNECT"
//     MethodOptions = "OPTIONS"
//     MethodTrace   = "TRACE"

type Service interface {
	Register(m *Mux) error
}

type ServiceFunc func(m *Mux) error

func (s ServiceFunc) Register(m *Mux) error {
	return s(m)
}
