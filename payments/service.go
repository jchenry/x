package payments

import (
	jch_http "github.com/jchenry/jchenry/http"
)

func Service(c Config) ServiceInstance {
	return ServiceInstance{c: c}
}

type ServiceInstance struct {
	c Config
}

func (si ServiceInstance) Register(uriBase string, s *jch_http.Server) {

}
