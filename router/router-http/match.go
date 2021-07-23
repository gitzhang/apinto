package router_http

import (
	"github.com/eolinker/goku-eosc/router"
	"github.com/eolinker/goku-eosc/service"
	"net/http"
)

type IMatcher interface {
	Match(req *http.Request) (service.IService,router.IEndpoint, bool)
}

type Matcher struct {
	r router.IRouter
	services map[string]service.IService
}

func (m *Matcher) Match(req *http.Request) (service.IService,router.IEndpoint, bool) {

	sources:=newHttpSources(req)
	endpoint,has:=m.r.Router(sources)
	if !has{
		return nil,nil,false
	}

	s,has:=m.services[endpoint.Target()]

	return s,endpoint,has
}


type  HttpSources struct {
	req * http.Request
}

func newHttpSources(req *http.Request) *HttpSources {
	return &HttpSources{req: req}
}

func (h *HttpSources) Get(cmd string) (string, bool) {

}
