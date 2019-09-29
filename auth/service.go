package auth

import (
	"net/http"

	"github.com/codegangsta/negroni"
	jch_http "github.com/jchenry/jchenry/http"
)

func Service(c Config) ServiceInstance {
	return ServiceInstance{c: c}
}

type ServiceInstance struct {
	c Config
}

func (si ServiceInstance) Register(uriBase string, s *jch_http.Server) {

	s.GET(uriBase+"/login", "login endpoint", http.HandlerFunc(NewLoginHandler(si.c)))
	s.GET(uriBase+"/logout", "logout endpoint", http.HandlerFunc(LogoutHandler))
	s.GET(uriBase+"/callback", "oidc callback", http.HandlerFunc(NewCallbackHandler(si.c)))
	s.GET(uriBase+"/user", "user info endpoint", negroni.New(
		negroni.HandlerFunc(IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(UserHandler)),
	))
}
