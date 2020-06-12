package auth

import (
	"net/http"

	"github.com/codegangsta/negroni"
	_http "github.com/jchenry/jchenry/internal/http"
	"gopkg.in/auth0.v1/management"
)

func Service(c Config) ServiceInstance {
	return ServiceInstance{c: c}
}

type ServiceInstance struct {
	c Config
}

func (si ServiceInstance) Register(uriBase string, s *_http.Server) {

	s.Get(uriBase+"/login", "login endpoint", http.HandlerFunc(NewLoginHandler(si.c)))
	s.Get(uriBase+"/logout", "logout endpoint", http.HandlerFunc(LogoutHandler))
	s.Get(uriBase+"/callback", "oidc callback", http.HandlerFunc(NewCallbackHandler(si.c)))
	s.Get(uriBase+"/user", "user info endpoint", negroni.New(
		negroni.HandlerFunc(IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(UserHandler)),
	))
}

func (si ServiceInstance) UpdateUser(u User) error {
	m, err := management.New(si.c.Domain, si.c.ManagementClientID, si.c.ManagementClientSecret)
	if err != nil {
		return err
	}

	um := management.NewUserManager(m)

	return um.Update(u.ID, &management.User{AppMetadata: u.Apps})
}
